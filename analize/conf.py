

TYPE_LOGS =  [
    # "fiber_ramp_up_deep", "fiber_ramp_up_flat", "gin_ramp_up_deep", "gin_ramp_up_flat",
    "fiber_high_deep", "fiber_high_flat", "gin_high_deep", "gin_high_flat",
] #  

GROUPS = {
    "ramp_up_deep": ["fiber_ramp_up_deep", "gin_ramp_up_deep"],
    "ramp_up_flat": ["fiber_ramp_up_flat", "gin_ramp_up_flat"],
    "high_deep": ["fiber_high_deep", "gin_high_deep"],
    "high_flat": ["fiber_high_flat", "gin_high_flat"],
    "spike_deep": ["fiber_spike_deep", "gin_spike_deep"],
    "spike_flat": ["fiber_spike_flat", "gin_spike_flat"],
}
DIRECTORY_METRICS = "./metrics_data"
