import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import os
from conf import DIRECTORY_METRICS, GROUPS
'''
1760781599.189	flat_request	2413	0	0	0	0	0	0	0	0	200
временная метка (в формате Unix Epoch (количество секунд с 1 января 1970 года)) .189 — это миллисекунды.
тип запроса
2413 - Время ответа в миллисекундах.
'''


window = 150

# return timestamp_response_dict[временная метка запроса отноистельно первого запроса] = длительности запросов пришедших в это время
def parseFile(filename: str, timestamp_responses_dict: dict):    
    try:
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
                    
                timestamp_ms = int(parts[0].replace(".", "")) - start_timestamp
                response_time_ms = int(parts[2])
                if timestamp_ms not in timestamp_responses_dict:
                    timestamp_responses_dict[timestamp_ms] = [response_time_ms]
                else:
                    timestamp_responses_dict[timestamp_ms].append(response_time_ms)         
    except (ValueError, IndexError) as e:
        print(f"Ошибка парсинга строки: {line}")

def parseFilesInDir(directory):
    timestamp_response_dict = dict()
    for subdir in os.listdir(directory):
        if not subdir.isnumeric():
            continue
        filepath = os.path.join(directory, subdir, "pandora.phout")
        print(f"Обрабатывается файл: {filepath}")
        
        if os.path.isfile(filepath):
            parseFile(filepath, timestamp_response_dict)
    
    for k, v in timestamp_response_dict.items():
        timestamp_response_dict[k] = sum(v) / len(v)

    sorted_times = sorted(timestamp_response_dict.keys())
    throughputs = [timestamp_response_dict[t] for t in sorted_times]
    
    return sorted_times, throughputs

def addGraph(time_starts, durations_ms, type, color1, color2):
    # Создаем линейный график
    plt.plot(time_starts, durations_ms, alpha=0.2, linewidth=1.5, color=color1, label=f"{type}")
    # Добавляем скользящее среднее для тренда
    if len(time_starts) > 10:
        window_size = min(window, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'duration': durations_ms})
        df = df.sort_values('time')
        df['moving_avg'] = df['duration'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color=color2, linewidth=2, 
                label=f'{type} скользящее среднее (окно={window_size})')


def plot(fiber_time_starts, fiber_durations_ms, gin_time_starts, gin_durations_ms, output_file, type):
    plt.figure(figsize=(14, 8))
    addGraph(fiber_time_starts, fiber_durations_ms, "fiber", 'blue', 'red')
    addGraph(gin_time_starts, gin_durations_ms, "gin", 'cyan', 'green')
    
    plt.xlabel('Время (мс)', fontsize=12)
    plt.ylabel('Скорость обработки (запросов/мкс)', fontsize=12)
    plt.title(f'Скорость обработки запросов во времени {type}', fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()
    
    # Добавляем статистику на график
    # stats_text = f'Всего временных точек: {len(time_starts)}\n'
    # stats_text += f'Макс. скорость: {max(throughputs):.0f} запр/сек\n'
    # stats_text += f'Ср. скорость: {np.mean(throughputs):.1f} запр/сек\n'
    # stats_text += f'Мин. скорость: {min(throughputs):.0f} запр/сек'
    # plt.annotate(stats_text, xy=(0.02, 0.98), xycoords='axes fraction',
    #             bbox=dict(boxstyle="round,pad=0.3", facecolor="white", alpha=0.8),
    #             verticalalignment='top', fontsize=10)
    # plt.tight_layout()
    
    plt.savefig(output_file, dpi=300, bbox_inches='tight')
    plt.close()
    print(f"График сохранен как {output_file}")

def main_throughput():
    try:
        for title, gr in GROUPS.items():
            fiber_time_starts, fiber_durations_ms = parseFilesInDir(f"{DIRECTORY_METRICS}/{gr[0]}")
            gin_time_starts, gin_durations_ms = parseFilesInDir(f"{DIRECTORY_METRICS}/{gr[1]}")
            os.makedirs(f"./img/{title}", exist_ok=True)
            plot(
                fiber_time_starts, fiber_durations_ms, 
                gin_time_starts, gin_durations_ms,
                f"./img/{title}/req_proc_plot.png", title)
            
    except FileNotFoundError:
        print(f"Директория не найдена")
    except Exception as e:
        print(f"Ошибка: {e}")
        import traceback
        traceback.print_exc()

if __name__ == "__main__":
    main_throughput()

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