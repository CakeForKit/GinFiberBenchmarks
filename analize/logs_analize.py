import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import os
from conf import DIRECTORY_METRICS, GROUPS

window = 100
LOGS_SHORT_FILENAME = "logs_time_series.txt"

def ParseFile(filename: str):
    """Парсинг файла с данными в формате timeStart duration"""
    try:
        time_duration_dict = dict()
        with open(filename, 'r') as f:
            f.readline()
            line = f.readline().strip()
            assert(line)
            parts = line.split()
            assert(len(parts) >= 2)

            prev_time_starts, sum_durations = float(parts[0]), float(parts[1])
            cnt_in_one = 1

            for line in f:
                line = line.strip()
                assert(line)
                parts = line.split()
                assert(len(parts) >= 2)

                time_start, duration = float(parts[0]), float(parts[1])
                if time_start == prev_time_starts:
                    cnt_in_one += 1
                    sum_durations += duration
                else:
                    time_duration_dict[prev_time_starts] = sum_durations / cnt_in_one
                    cnt_in_one = 1
                    prev_time_starts, sum_durations = time_start, duration
    except FileNotFoundError:
        print(f"Файл {filename} не найден")
    return time_duration_dict


def ParseFilesInDir(directory):
    med_time_duration_dict = dict()
    for subdir in os.listdir(directory):
        if not subdir.isnumeric():
            continue
        filepath = os.path.join(directory, subdir, LOGS_SHORT_FILENAME)
        print(f"Обрабатывается файл: {filepath}")

        if os.path.isfile(filepath):
            time_duration_dict = ParseFile(filepath)
            for k, v in time_duration_dict.items():
                if k not in med_time_duration_dict:
                    med_time_duration_dict[k] = [v]
                else:
                    med_time_duration_dict[k].append(v)

    for k, v in med_time_duration_dict.items():
        med_time_duration_dict[k] = sum(v) / len(v)
    
    time_starts = list(sorted(med_time_duration_dict.keys()))
    durations_ms = [med_time_duration_dict[i] for i in time_starts]

    return time_starts, durations_ms
            
def AddGraph(time_starts, durations_ms, type, color1, color2):
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

def Plot(fiber_time_starts, fiber_durations_ms, gin_time_starts, gin_durations_ms, output_file, type):
    """Построение графика timeStart vs duration"""
    plt.figure(figsize=(12, 8))
    AddGraph(fiber_time_starts, fiber_durations_ms, "fiber", 'blue', 'red')
    AddGraph(gin_time_starts, gin_durations_ms, "gin", 'cyan', 'green')
    
    plt.xlabel('Время начала (мс)', fontsize=12)
    plt.ylabel('Длительность сериализации (мкс)', fontsize=12)
    plt.title(f'Время сериализации {type}', fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()

    plt.savefig(output_file, dpi=300, bbox_inches='tight')
    plt.close()
    print(f"График сохранен как {output_file}")
    

def main_logs():
    try:
        for title, gr in GROUPS.items():
            fiber_time_starts, fiber_durations_ms = ParseFilesInDir(f"{DIRECTORY_METRICS}/{gr[0]}")
            gin_time_starts, gin_durations_ms = ParseFilesInDir(f"{DIRECTORY_METRICS}/{gr[1]}")
            os.makedirs(f"./img/{title}", exist_ok=True)
            Plot(
                fiber_time_starts, fiber_durations_ms, 
                gin_time_starts, gin_durations_ms,
                f"./img/{title}/time_series_plot.png", title)
        
    except FileNotFoundError:
        print(f"Файл ... не найден")
    except Exception as e:
        print(f"Ошибка: {e}")
        import traceback
        traceback.print_exc()

if __name__ == "__main__":
    main_logs()

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
    