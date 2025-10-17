import matplotlib.pyplot as plt
import pandas as pd
from datetime import datetime, timedelta
import numpy as np

filename = "./metrics_data/logs/flat_20251017_163859"  # Замените на путь к вашему файлу

def parse_data_file(filename: str):
    """Парсинг файла с данными в формате timeStart duration"""
    time_starts = []
    durations_ms = []
    
    with open(filename, 'r') as f:
        
        print(f.readline())
        for line in f:
            line = line.strip()
            if not line:
                continue
                
            parts = line.split()
            if len(parts) >= 2:
                time_start_str = parts[0]
                duration_str = parts[1]
                
                # Парсим timeStart (предполагаем, что это секунды от начала)
                print(time_start_str)
                time_start_sec = float(time_start_str)
                
                # Парсим duration и конвертируем в миллисекунды
                if 'µs' in parts[1]:
                    duration_ms = float(duration_str) / 1000  # микросекунды в миллисекунды
                else:
                    duration_ms = float(duration_str)  # уже в миллисекундах
                
                time_starts.append(time_start_sec * 1000)  # конвертируем в миллисекунды
                durations_ms.append(duration_ms)
    
    return time_starts, durations_ms

def plot_time_series(time_starts, durations_ms, output_file):
    """Построение графика timeStart vs duration"""
    plt.figure(figsize=(12, 8))
    
    # Создаем scatter plot
    plt.scatter(time_starts, durations_ms, alpha=0.6, s=20, color='blue', label='Измерения')
    
    # Добавляем скользящее среднее для тренда
    if len(time_starts) > 10:
        window_size = min(50, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'duration': durations_ms})
        df = df.sort_values('time')
        df['moving_avg'] = df['duration'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color='red', linewidth=2, 
                label=f'Скользящее среднее (окно={window_size})')
    
    plt.xlabel('Время начала (ms)', fontsize=12)
    plt.ylabel('Длительность сериализации (ms)', fontsize=12)
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

def plot_histogram(durations_ms, output_file):
    """Построение гистограммы распределения длительностей"""
    plt.figure(figsize=(10, 6))
    
    plt.hist(durations_ms, bins=50, alpha=0.7, color='green', edgecolor='black')
    plt.xlabel('Длительность сериализации (ms)', fontsize=12)
    plt.ylabel('Количество измерений', fontsize=12)
    plt.title('Распределение длительностей сериализации', fontsize=14)
    plt.grid(True, alpha=0.3)
    
    # Добавляем вертикальные линии для перцентилей
    percentiles = [50, 75, 90, 95, 99]
    colors = ['red', 'orange', 'yellow', 'purple', 'brown']
    
    for p, color in zip(percentiles, colors):
        value = np.percentile(durations_ms, p)
        plt.axvline(value, color=color, linestyle='--', alpha=0.8, 
                   label=f'P{p}: {value:.3f}ms')
    
    plt.legend()
    plt.tight_layout()
    
    if output_file:
        plt.savefig(output_file, dpi=300, bbox_inches='tight')
        print(f"Гистограмма сохранена как {output_file}")
    
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
        time_starts, durations_ms = parse_data_file(filename)
        
        # Вывод статистики
        print_statistics(time_starts, durations_ms)
        
        # Построение графиков
        plot_time_series(time_starts, durations_ms, "time_series_plot.png")
        plot_histogram(durations_ms, "duration_histogram.png")
        
    except FileNotFoundError:
        print(f"Файл {filename} не найден")
    except Exception as e:
        print(f"Ошибка: {e}")
        import traceback
        traceback.print_exc()

# Альтернативная версия для данных в другом формате
def parse_data_file_alternative(filename: str):
    """Альтернативный парсер для разных форматов данных"""
    time_starts = []
    durations_ms = []
    
    with open(filename, 'r') as f:
        for line_num, line in enumerate(f, 1):
            line = line.strip()
            if not line:
                continue
                
            try:
                # Пробуем разные разделители
                for separator in [' ', '\t', ',', ';']:
                    if separator in line:
                        parts = line.split(separator)
                        if len(parts) >= 2:
                            # Убираем все не-цифровые символы кроме точки и минуса
                            time_str = ''.join(c for c in parts[0] if c.isdigit() or c in '.-')
                            duration_str = ''.join(c for c in parts[1] if c.isdigit() or c in '.-')
                            
                            time_start = float(time_str)
                            duration = float(duration_str)
                            
                            # Автоопределение единиц измерения
                            if 'µs' in line or 'us' in line or duration < 0.1:
                                duration_ms = duration / 1000
                            else:
                                duration_ms = duration
                            
                            time_starts.append(time_start * 1000)  # в миллисекунды
                            durations_ms.append(duration_ms)
                            break
            except ValueError as e:
                print(f"Пропущена строка {line_num}: {line} (ошибка: {e})")
                continue
    
    return time_starts, durations_ms

if __name__ == "__main__":
    main()
    