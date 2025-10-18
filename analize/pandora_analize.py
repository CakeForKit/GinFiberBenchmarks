import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import os
from datetime import datetime

'''
1760781599.189	flat_request	2413	0	0	0	0	0	0	0	0	200
временная метка (в формате Unix Epoch (количество секунд с 1 января 1970 года)) .189 — это миллисекунды.
тип запроса
2413 - Время ответа в миллисекундах.
'''

directoryLogs = "./metrics_data/save/"

def parse_phout_files_in_dir(directory):
    """
    Парсинг всех phout файлов в директории 
    """
    timestamp_response_dict = dict()
    
    for subdir in os.listdir(directory):
        
        # if subdir not in ['1', '2']:
        #     continue
        # print(subdir)
        filepath = os.path.join(directory, subdir, "flat_results.phout")
        print(f"Обрабатывается файл: {filepath}")
        
        if os.path.isfile(filepath):
            parse_phout_file(filepath, timestamp_response_dict)
    
    
    for k, v in timestamp_response_dict.items():
        timestamp_response_dict[k] = sum(v) / len(v)

    # Сортируем по времени и преобразуем в списки
    sorted_times = sorted(timestamp_response_dict.keys())
    throughputs = [timestamp_response_dict[t] for t in sorted_times]
    
    print(f"Временной диапазон: [{min(sorted_times)}, {max(sorted_times)}]")
    print(f"Всего временных точек: {len(sorted_times)}")
    
    return sorted_times, throughputs

# return timestamp_response_dict[временная метка запроса отноистельно первого запроса] = длительности запросов пришедших в это время
def parse_phout_file(filename: str, timestamp_responses_dict: dict):
    """
    Парсинг одного phout файла
    timestamp_response_dict[временная метка запроса отноистельно первого запроса] = длительности запросов пришедших в это время
    """
    
    with open(filename, 'r') as f:
        line = f.readline().strip()
        assert(line)
        parts = line.split('\t')
        assert(len(parts) >= 3)
        start_timestamp, response_time_ms = int(parts[0].replace(".", "")), int(parts[2])
        timestamp_responses_dict[0] = [response_time_ms]

        for line in f:
            line = line.strip()
            assert(line)
                
            parts = line.split('\t')
            assert(len(parts) >= 3)
                
            try:
                timestamp_ms = int(parts[0].replace(".", "")) - start_timestamp
                response_time_ms = int(parts[2])
                if timestamp_ms not in timestamp_responses_dict:
                    timestamp_responses_dict[timestamp_ms] = [response_time_ms]
                else:
                    timestamp_responses_dict[timestamp_ms].append(response_time_ms)
                    
            except (ValueError, IndexError) as e:
                print(f"Ошибка парсинга строки: {line}")
                continue

def plot_throughput_time_series(time_starts, throughputs, output_file=None):
    """Построение графика времени начала vs скорость обработки запросов"""
    plt.figure(figsize=(14, 8))
    
    # Строим линейный график
    plt.plot(time_starts, throughputs, alpha=0.2, linewidth=1.5, color='blue', label='Скорость обработки')
    
    # Добавляем скользящее среднее для сглаживания
    if len(time_starts) > 10:
        window_size = min(50, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'throughput': throughputs})
        df['moving_avg'] = df['throughput'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color='red', linewidth=2, 
                label=f'Скользящее среднее (окно={window_size})')
    
    plt.xlabel('Время (секунды)', fontsize=12)
    plt.ylabel('Скорость обработки (запросов/секунду)', fontsize=12)
    plt.title('Скорость обработки запросов во времени', fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()
    
    # Добавляем статистику на график
    stats_text = f'Всего временных точек: {len(time_starts)}\n'
    stats_text += f'Макс. скорость: {max(throughputs):.0f} запр/сек\n'
    stats_text += f'Ср. скорость: {np.mean(throughputs):.1f} запр/сек\n'
    stats_text += f'Мин. скорость: {min(throughputs):.0f} запр/сек'
    
    plt.annotate(stats_text, xy=(0.02, 0.98), xycoords='axes fraction',
                bbox=dict(boxstyle="round,pad=0.3", facecolor="white", alpha=0.8),
                verticalalignment='top', fontsize=10)
    
    plt.tight_layout()
    
    if output_file:
        plt.savefig(output_file, dpi=300, bbox_inches='tight')
        print(f"График сохранен как {output_file}")
    
    plt.show()

def print_throughput_statistics(time_starts, throughputs):
    """Вывод статистики по скорости обработки"""
    print("=== СТАТИСТИКА СКОРОСТИ ОБРАБОТКИ ===")
    print(f"Всего временных интервалов: {len(time_starts)}")
    print(f"Общее время наблюдения: {max(time_starts) - min(time_starts):.0f} секунд")
    print(f"Первый замер: {time_starts[0]}")
    print(f"Последний замер: {time_starts[-1]}")
    print()
    
    print("=== СТАТИСТИКА ПРОПУСКНОЙ СПОСОБНОСТИ ===")
    print(f"Минимальная скорость: {min(throughputs):.0f} запр/сек")
    print(f"Максимальная скорость: {max(throughputs):.0f} запр/сек")
    print(f"Средняя скорость: {np.mean(throughputs):.1f} запр/сек")
    print(f"Медианная скорость: {np.median(throughputs):.1f} запр/сек")
    print(f"Стандартное отклонение: {np.std(throughputs):.1f} запр/сек")
    print(f"Общее количество запросов: {sum(throughputs):.0f}")
    print()
    
    print("=== ПЕРЦЕНТИЛИ СКОРОСТИ ===")
    percentiles = [50, 75, 90, 95, 99]
    for p in percentiles:
        value = np.percentile(throughputs, p)
        print(f"P{p}: {value:.1f} запр/сек")

def main_throughput():
    """Основная функция для анализа скорости обработки"""
    try:
        # Чтение и парсинг данных
        time_starts, throughputs = parse_phout_files_in_dir(directoryLogs)
        
        if not time_starts:
            print("Не найдено данных для построения графика")
            return
        
        # Вывод статистики
        # print_throughput_statistics(time_starts, throughputs)
        
        # Построение графиков
        plot_throughput_time_series(time_starts, throughputs, "throughput_time_series.png")
        
    except FileNotFoundError:
        print(f"Директория {directoryLogs} не найдена")
    except Exception as e:
        print(f"Ошибка: {e}")
        import traceback
        traceback.print_exc()

if __name__ == "__main__":
    main_throughput()