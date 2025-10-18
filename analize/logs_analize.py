import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import os

# filename = "./metrics_data/logs/flat_20251017_163859"  # Замените на путь к вашему файлу
directoryLogs = "./metrics_data/save/"

def parse_files_in_dir(directory):
    med_time_duration_dict = dict()
    for filename in os.listdir(directory):
        filepath = os.path.join(directory, filename, "flat_logs.txt")
        print(filepath)
        if os.path.isfile(filepath):
            time_duration_dict = parse_data_file(filepath)
            for k, v in time_duration_dict.items():
                if k not in med_time_duration_dict:
                    med_time_duration_dict[k] = [v]
                else:
                    med_time_duration_dict[k].append(v)
    for k, v in med_time_duration_dict.items():
        med_time_duration_dict[k] = sum(med_time_duration_dict[k]) / len(med_time_duration_dict[k])
    
    time_starts = list(sorted(med_time_duration_dict.keys()))
    durations_ms = []
    for ts in time_starts:
        durations_ms.append(med_time_duration_dict[ts])
    print(f"time_starts: [{min(time_starts)}, {max(time_starts)}]")
    return time_starts, durations_ms
            

def parse_data_file(filename: str):
    """Парсинг файла с данными в формате timeStart duration"""
    print(f"filename: {filename}")
    lines = 0
    time_duration_dict = dict()
    with open(filename, 'r') as f:
        f.readline()
        line = f.readline().strip()
        assert(line)
        parts = line.split()
        assert(len(parts) >= 2)
        last_time_starts, last_durations = float(parts[0]), float(parts[1])
        cnt_in_one = 1
        for line in f:
            line = line.strip()
            assert(line)

            parts = line.split()
            assert(len(parts) >= 2)

            time_start, duration = float(parts[0]), float(parts[1])
            if time_start == last_time_starts:
                cnt_in_one += 1
                last_durations += duration
            else:
                time_duration_dict[last_time_starts] = last_durations / cnt_in_one
                cnt_in_one = 1
                last_time_starts, last_durations = time_start, duration
            lines += 1
    print(f"lines = {lines}")
    print(f"{min(time_duration_dict.keys())} - {max(time_duration_dict.keys())}")
    return time_duration_dict

def plot_time_series(time_starts, durations_ms, output_file):
    """Построение графика timeStart vs duration"""
    plt.figure(figsize=(12, 8))
    # Создаем scatter plot
    plt.plot(time_starts, durations_ms, alpha=0.2, linewidth=1.5, color='blue', label='Измерения')
    
    # Добавляем скользящее среднее для тренда
    if len(time_starts) > 10:
        window_size = min(50, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'duration': durations_ms})
        df = df.sort_values('time')
        df['moving_avg'] = df['duration'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color='red', linewidth=2, 
                label=f'Скользящее среднее (окно={window_size})')
    
    plt.xlabel('Время начала (ms)', fontsize=12)
    plt.ylabel('Длительность сериализации (us)', fontsize=12)
    plt.title('Время сериализации vs Время начала операции', fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()
    
    # Добавляем статистику на график
    stats_text = f'Всего измерений: {len(time_starts)}\n'
    stats_text += f'Макс. длительность: {max(durations_ms):.3f} ms\n'
    stats_text += f'Ср. длительность: {np.mean(durations_ms):.3f} ms'
    
    plt.annotate(stats_text, xy=(0.02, 0.98), xycoords='axes fraction',
                bbox=dict(boxstyle="round,pad=0.3", facecolor="white", alpha=0.8),
                verticalalignment='top', fontsize=10)
    
    plt.tight_layout()
    
    if output_file:
        plt.savefig(output_file, dpi=300, bbox_inches='tight')
        print(f"График сохранен как {output_file}")
    
    plt.show()

def print_statistics(time_starts, durations_ms):
    """Вывод статистики"""
    print("=== СТАТИСТИКА ДАННЫХ ===")
    print(f"Всего измерений: {len(time_starts)}")
    print(f"Общее время наблюдения: {max(time_starts) - min(time_starts):.2f} ms")
    print(f"Первый замер: {min(time_starts):.2f} ms")
    print(f"Последний замер: {max(time_starts):.2f} ms")
    print()
    
    print("=== СТАТИСТИКА ДЛИТЕЛЬНОСТЕЙ ===")
    print(f"Минимальная длительность: {min(durations_ms):.6f} ms")
    print(f"Максимальная длительность: {max(durations_ms):.6f} ms")
    print(f"Средняя длительность: {np.mean(durations_ms):.6f} ms")
    print(f"Медианная длительность: {np.median(durations_ms):.6f} ms")
    print(f"Стандартное отклонение: {np.std(durations_ms):.6f} ms")
    print()
    
    print("=== ПЕРЦЕНТИЛИ ===")
    percentiles = [50, 75, 90, 95, 99]
    for p in percentiles:
        value = np.percentile(durations_ms, p)
        print(f"P{p}: {value:.6f} ms")

def main():
    try:
        # Чтение и парсинг данных
        time_starts, durations_ms = parse_files_in_dir(directoryLogs)
        # print(time_starts)
        
        # Вывод статистики
        # print_statistics(time_starts, durations_ms)
        
        # Построение графиков
        plot_time_series(time_starts, durations_ms, "time_series_plot.png")
        # plot_histogram(durations_ms, "duration_histogram.png")
        
    except FileNotFoundError:
        print(f"Файл {directoryLogs}... не найден")
    except Exception as e:
        print(f"Ошибка: {e}")
        import traceback
        traceback.print_exc()

if __name__ == "__main__":
    main()
    