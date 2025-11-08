export interface HealthData {
  status: string;
  timestamp: number;
  uptime: string;
  version: {
    commit: string;
    go_version: string;
    build_time: string;
    build_mode: string;
  };
  runtime: {
    go_os: string;
    go_arch: string;
    num_goroutine: number;
    num_cpu: number;
  };
  memory: {
    alloc: number;
    total_alloc: number;
    sys: number;
    num_gc: number;
  };
}
