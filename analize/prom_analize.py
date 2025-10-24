from typing import List, Dict
import os
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
from conf import DIRECTORY_METRICS, GROUPS

class MetricParams:
    def __init__(self, graph_label: str, outputFilename: str):
        self.graph_label = graph_label
        self.outputFilename = outputFilename 


FILENAMES = {
    "container_cpu_usage_seconds_total.txt": MetricParams("Загрузка процессора в процентах", "cpu_usage_seconds_total.png"),
    "container_memory_usage_bytes.txt": MetricParams("Потребление памяти контейнером  (байты)", "memory_usage_bytes.png"),
    "goroutines_count.txt": MetricParams("Количество горутин", "goroutines_count.png"),
    "memory_allocations_bytes_heap.txt": MetricParams("Память выделенных объектов в куче (байты)", "memory_allocations_bytes_heap.png"),
    "memory_allocations_bytes_stack.txt": MetricParams("Память в стеках горутин (байты)", "memory_allocations_bytes_stack.png"),
    "total_http_request_counter.txt": MetricParams("Количество обработанных HTTP-запросов", "total_http_request_counter.png")
}
window = 10



def ParseFile(filename: str, res_list: List[List[float]]):
    ir = 0
    try:
        with open(filename, 'r') as f:
            line = f.readline().strip()
            assert(line)
            parts = line.split()
            assert(len(parts) == 2)
            start_timestamp, metric = int(parts[0]), float(parts[1])
            if ir >= len(res_list):
                res_list.append([])
            res_list[ir].append(metric)
            ir += 1

            for line in f:
                parts = line.strip().split()
                assert(len(parts) == 2)

                timestamp, metric = int(parts[0]), float(parts[1])
                if ir >= len(res_list):
                    res_list.append([])
                res_list[ir].append(metric)
                ir += 1
        # print(f"ir = {ir}")
    except (ValueError, IndexError) as e:
        print(f"Ошибка парсинга строки: {line} {e}")
        raise
    
def ParseFilesInDir(directory: str, shortFilename: str) -> List[float]:
    res_list = list()
    for subdir in os.listdir(directory):
        filepath = os.path.join(directory, subdir, shortFilename)
        # if int(subdir) <= 2:
        #     continue
        print(f"Обрабатывается файл: {filepath}")

        if os.path.isfile(filepath):
            ParseFile(filepath, res_list)
        print("res_list: ", len(res_list))
    
    # x = timestamp from 0, y = metric
    med_res = list()
    for lst in res_list:
        med_res.append(sum(lst) / len(lst))

    # Вычисляем статистику для всего набора данных
    all_metrics = [metric for sublist in res_list for metric in sublist]
    if all_metrics:
        stats = {
            'min': min(all_metrics),
            'max': max(all_metrics),
            'median': np.median(all_metrics),
            'mean': np.mean(all_metrics),
            'std': np.std(all_metrics)
        }
    else:
        stats = {'min': 0, 'max': 0, 'median': 0, 'mean': 0, 'std': 0}
    
    return med_res, stats

def format_metric_value(value: float, metric_name: str) -> str:
    """Форматирование значений метрик для читаемости"""
    if "memory" in metric_name.lower() or "bytes" in metric_name.lower():
        if value >= 1024**3:  # GB
            return f"{value/1024**3:.2f} GB"
        elif value >= 1024**2:  # MB
            return f"{value/1024**2:.2f} MB"
        elif value >= 1024:  # KB
            return f"{value/1024:.2f} KB"
        else:
            return f"{value:.0f} B"
    elif "cpu" in metric_name.lower():
        return f"{value:.2f}%"
    else:
        return f"{value:.2f}"

def AddGraph(time_starts, durations_ms, type, color1, color2):
    # Создаем линейный график
    # plt.plot(time_starts, durations_ms, alpha=0.2, linewidth=1.5, color=color1, label=f"{type}")
    # Добавляем скользящее среднее для тренда
    if len(time_starts) > 10:
        window_size = min(window, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'duration': durations_ms})
        df = df.sort_values('time')
        df['moving_avg'] = df['duration'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color=color2, linewidth=2, 
                label=f'{type} скользящее среднее (окно={window_size})')

def add_statistics_to_plot(fiber_stats: Dict[str, Dict[str, float]], 
                          gin_stats: Dict[str, Dict[str, float]], 
                          graph_label: str):
    """Добавление статистической информации на график"""
    
    # Создаем текстовую строку со статистикой
    stats_text = "Статистика по компонентам:\n\n"
    
    for component in fiber_stats.keys():
        stats_text += f"{component.upper()}:\n"
        
        if "memory" in component or "bytes" in component:
            # Для метрик памяти
            fiber_min = format_metric_value(fiber_stats[component]['min'], component)
            fiber_max = format_metric_value(fiber_stats[component]['max'], component)
            fiber_med = format_metric_value(fiber_stats[component]['median'], component)
            
            gin_min = format_metric_value(gin_stats[component]['min'], component)
            gin_max = format_metric_value(gin_stats[component]['max'], component)
            gin_med = format_metric_value(gin_stats[component]['median'], component)
            
            stats_text += f"  Fiber - min: {fiber_min}, max: {fiber_max}, med: {fiber_med}\n"
            stats_text += f"  Gin   - min: {gin_min}, max: {gin_max}, med: {gin_med}\n\n"
            
        elif "cpu" in component:
            # Для CPU
            fiber_min = f"{fiber_stats[component]['min']:.2f}%"
            fiber_max = f"{fiber_stats[component]['max']:.2f}%"
            fiber_med = f"{fiber_stats[component]['median']:.2f}%"
            
            gin_min = f"{gin_stats[component]['min']:.2f}%"
            gin_max = f"{gin_stats[component]['max']:.2f}%"
            gin_med = f"{gin_stats[component]['median']:.2f}%"
            
            stats_text += f"  Fiber - min: {fiber_min}, max: {fiber_max}, med: {fiber_med}\n"
            stats_text += f"  Gin   - min: {gin_min}, max: {gin_max}, med: {gin_med}\n\n"
            
        else:
            # Для остальных метрик
            fiber_min = f"{fiber_stats[component]['min']:.2f}"
            fiber_max = f"{fiber_stats[component]['max']:.2f}"
            fiber_med = f"{fiber_stats[component]['median']:.2f}"
            
            gin_min = f"{gin_stats[component]['min']:.2f}"
            gin_max = f"{gin_stats[component]['max']:.2f}"
            gin_med = f"{gin_stats[component]['median']:.2f}"
            
            stats_text += f"  Fiber - min: {fiber_min}, max: {fiber_max}, med: {fiber_med}\n"
            stats_text += f"  Gin   - min: {gin_min}, max: {gin_max}, med: {gin_med}\n\n"
    
    # Добавляем текст на график
    plt.figtext(0.02, 0.02, stats_text, fontsize=10, 
                bbox=dict(boxstyle="round,pad=0.3", facecolor="lightgray", alpha=0.8),
                verticalalignment='bottom')

def Plot(
        fiber_med_res: List[float], gin_med_res: List[float], 
        fiber_stats: Dict[str, float], gin_stats: Dict[str, float],
        graph_label: str, outputFilename, metric_name: str,
        ):
    fiber_timestamps = [i for i in range(len(fiber_med_res))]
    fiber_metrics = [fiber_med_res[i] for i in fiber_timestamps]
    gin_timestamps = [i for i in range(len(gin_med_res))]
    gin_metrics = [gin_med_res[i] for i in gin_timestamps]

    plt.figure(figsize=(12, 8))
    AddGraph(fiber_timestamps, fiber_metrics, "fiber", 'blue', 'red')
    AddGraph(gin_timestamps, gin_metrics, "gin", 'cyan', 'green')
    
    plt.xlabel('Время (секунды)', fontsize=10)
    plt.ylabel(graph_label, fontsize=10)
    plt.title(graph_label, fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()

    # Добавляем статистику на график
    stats_data = {
        metric_name: {
            'fiber': fiber_stats,
            'gin': gin_stats
        }
    }
    add_statistics_to_plot(
        {metric_name: fiber_stats}, 
        {metric_name: gin_stats}, 
        graph_label
    )
    # Регулируем layout чтобы освободить место для статистики
    plt.tight_layout(rect=[0, 0.2, 1, 0.95])

    plt.savefig(outputFilename, dpi=300, bbox_inches='tight')
    plt.close()
    print(f"График сохранен как {outputFilename}")

def main_prom():
    all_stats = {}
    for shortFilename, params in FILENAMES.items():
        for type, gr in GROUPS.items():
            if type not in all_stats:
                all_stats[type] = {'fiber': {}, 'gin': {}}

            fiber_med_res, fiber_stats = ParseFilesInDir(directory=f"{DIRECTORY_METRICS}/{gr[0]}", shortFilename=shortFilename)
            gin_med_res, gin_stats = ParseFilesInDir(directory=f"{DIRECTORY_METRICS}/{gr[1]}", shortFilename=shortFilename)
            # Сохраняем статистику для текущей метрики
            metric_key = shortFilename.replace('.txt', '')
            all_stats[type]['fiber'][metric_key] = fiber_stats
            all_stats[type]['gin'][metric_key] = gin_stats

            os.makedirs(f"./img/{type}", exist_ok=True)
            Plot(
                fiber_med_res, 
                gin_med_res,
                fiber_stats,
                gin_stats,
                params.graph_label + " " + type, f"./img/{type}/{params.outputFilename}",
                metric_key)
            

if __name__ == '__main__':
    main_prom()