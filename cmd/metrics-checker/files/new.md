# HELP base_memory_committedHeap_bytes Displays the amount of memory that is committed for the Java virtual machine to use.
# TYPE base_memory_committedHeap_bytes gauge
base_memory_committedHeap_bytes 3.9891034E-316
# HELP base_memory_initNonHeap_bytes Displays the initial amount of allocated memory in bytes for non-heap.
# TYPE base_memory_initNonHeap_bytes gauge
base_memory_initNonHeap_bytes 3.788353E-317
# HELP base_cpu_processCpuTime Displays the CPU time used by the process on which the Java virtual machine is running in nanoseconds
# TYPE base_cpu_processCpuTime gauge
base_cpu_processCpuTime 5.785546262E-314
# HELP base_thread_totalStarted Total number of started threads
# TYPE base_thread_totalStarted gauge
base_thread_totalStarted 50.0
# HELP base_memory_maxHeap_bytes Displays the maximum amount of memory in bytes that can be used for memory management.
# TYPE base_memory_maxHeap_bytes gauge
base_memory_maxHeap_bytes 2.652494739E-315
# HELP base_cpu_availableProcessors Displays the number of processors available to the Java virtual machine. This value may change during a particular invocation of the virtual machine.
# TYPE base_cpu_availableProcessors gauge
base_cpu_availableProcessors 12.0
# HELP base_classloader_unloadedClasses_total Displays the total number of classes unloaded since the Java virtual machine has started execution.
# TYPE base_classloader_unloadedClasses_total counter
base_classloader_unloadedClasses_total 0.0
# HELP base_thread_max_count Displays the peak live thread count since the Java virtual machine started or peak was reset. This includes daemon and non-daemon threads.
# TYPE base_thread_max_count gauge
base_thread_max_count 47.0
# HELP base_memory_committedNonHeap_bytes Displays the amount of memory that is committed for the Java virtual machine to use.
# TYPE base_memory_committedNonHeap_bytes gauge
base_memory_committedNonHeap_bytes 4.37117663E-316
# HELP base_classloader_loadedClasses_total Displays the total number of classes that have been loaded since the Java virtual machine has started execution.
# TYPE base_classloader_loadedClasses_total counter
base_classloader_loadedClasses_total 10790.0
# HELP base_cpu_systemLoadAverage Displays the system load average for the last minute. The system load average is the sum of the number of runnable entities queued to the available processors and the number of runnable entities running on the available processors averaged over a period of time. The way in which the load average is calculated is operating system specific but is typically a damped time-dependent average. If the load average is not available, a negative value is displayed. This attribute is designed to provide a hint about the system load and may be queried frequently. The load average may be unavailable on some platform where it is expensive to implement this method.
# TYPE base_cpu_systemLoadAverage gauge
base_cpu_systemLoadAverage 0.89013671875
# HELP base_memory_initHeap_bytes Displays the initial amount of allocated heap memory in bytes.
# TYPE base_memory_initHeap_bytes gauge
base_memory_initHeap_bytes 3.31561842E-316
# HELP base_gc_total Displays the total number of collections that have occurred. This attribute lists -1 if the collection count is undefined for this collector.
# TYPE base_gc_total counter
base_gc_total{name="G1 Young Generation",} 14.0
base_gc_total{name="G1 Old Generation",} 0.0
# HELP base_thread_count Current Thread count
# TYPE base_thread_count gauge
base_thread_count 46.0
# HELP base_cpu_processCpuLoad Displays the "recent cpu usage" for the Java Virtual Machine process.
# TYPE base_cpu_processCpuLoad gauge
base_cpu_processCpuLoad 0.001245385568236909
# HELP base_jvm_uptime Displays the uptime of the Java virtual machine
# TYPE base_jvm_uptime gauge
base_jvm_uptime 229179.0
# HELP base_classloader_loadedClasses_count Displays the number of classes that are currently loaded in the Java virtual machine.
# TYPE base_classloader_loadedClasses_count gauge
base_classloader_loadedClasses_count 10790.0
# HELP base_thread_daemon_count Displays the current number of live daemon threads.
# TYPE base_thread_daemon_count gauge
base_thread_daemon_count 37.0
# HELP base_gc_time_total Displays the approximate accumulated collection elapsed time in milliseconds. This attribute displays -1 if the collection elapsed time is undefined for this collector. The Java virtual machine implementation may use a high resolution timer to measure the elapsed time. This attribute may display the same value even if the collection count has been incremented if the collection elapsed time is very short.
# TYPE base_gc_time_total counter
base_gc_time_total{name="G1 Young Generation",} 63.0
base_gc_time_total{name="G1 Old Generation",} 0.0
# HELP base_memory_usedNonHeap_bytes Displays the amount of used memory.
# TYPE base_memory_usedNonHeap_bytes gauge
base_memory_usedNonHeap_bytes 4.1539518E-316
# HELP base_memory_maxNonHeap_bytes Displays the maximum amount of memory in bytes that can be used for memory management.
# TYPE base_memory_maxNonHeap_bytes gauge
base_memory_maxNonHeap_bytes NaN
# HELP base_memory_usedHeap_bytes Displays the amount of used memory.
# TYPE base_memory_usedHeap_bytes gauge
base_memory_usedHeap_bytes 2.06092883E-316
# HELP vendor_cache_manager_default_number_of_cache_configurations The total number of defined cache configurations.
# TYPE vendor_cache_manager_default_number_of_cache_configurations gauge
vendor_cache_manager_default_number_of_cache_configurations{node="fercoli-mac-20931",} 11.0
# HELP vendor_cache_manager_default_cache_container_stats_hits Cache container total number of cache attribute hits
# TYPE vendor_cache_manager_default_cache_container_stats_hits gauge
vendor_cache_manager_default_cache_container_stats_hits{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_timestamper Next seqno issued by the timestamper
# TYPE vendor_jgroups_cluster_unicast3_get_timestamper gauge
vendor_jgroups_cluster_unicast3_get_timestamper{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_num_purges Number of purges in all (receive and send) windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_num_purges gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_num_purges{node="fercoli-mac-20931",} 0.0
# HELP vendor_cpu_processCpuLoad Displays the "recent cpu usage" for the Java Virtual Machine process.
# TYPE vendor_cpu_processCpuLoad gauge
vendor_cpu_processCpuLoad 0.21497584541062803
# HELP vendor_cache_manager_default_cache_container_health_number_of_nodes Total nodes in the cluster
# TYPE vendor_cache_manager_default_cache_container_health_number_of_nodes gauge
vendor_cache_manager_default_cache_container_health_number_of_nodes{node="fercoli-mac-20931",} 1.0
# HELP vendor_cache_manager_default_cache_container_stats_average_write_time Cache container average number of milliseconds for all write operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_write_time gauge
vendor_cache_manager_default_cache_container_stats_average_write_time{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_average_remove_time_nanos Cache container total average number of nanoseconds for all remove operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_remove_time_nanos gauge
vendor_cache_manager_default_cache_container_stats_average_remove_time_nanos{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_time_since_reset Number of seconds since the cache container statistics were last reset
# TYPE vendor_cache_manager_default_cache_container_stats_time_since_reset gauge
vendor_cache_manager_default_cache_container_stats_time_since_reset{node="fercoli-mac-20931",} 224.0
# HELP vendor_cache_manager_default_number_of_running_caches The total number of running caches, including the default cache.
# TYPE vendor_cache_manager_default_number_of_running_caches gauge
vendor_cache_manager_default_number_of_running_caches{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_off_heap_memory_used Amount in bytes of off-heap memory used by this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_off_heap_memory_used gauge
vendor_cache_manager_default_cache_container_stats_off_heap_memory_used{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_thread_pool_size Current number of threads in the thread pool
# TYPE vendor_jgroups_cluster_tcp_get_thread_pool_size gauge
vendor_jgroups_cluster_tcp_get_thread_pool_size{node="fercoli-mac-20931",} 2.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_capacity Capacity of the retransmit buffer. Computed as xmit_table_num_rows * xmit_table_msgs_per_row
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_capacity gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_capacity{node="fercoli-mac-20931",} 51200.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_num_moves Number of retransmit table moves
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_num_moves gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_num_moves{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_receive_connections Returns the number of incoming (receive) connections
# TYPE vendor_jgroups_cluster_unicast3_get_num_receive_connections gauge
vendor_jgroups_cluster_unicast3_get_num_receive_connections{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cluster_container_stats_memory_used The amount of memory used by JVMs across the cluster in bytes
# TYPE vendor_cache_manager_default_cluster_container_stats_memory_used gauge
vendor_cache_manager_default_cluster_container_stats_memory_used{node="fercoli-mac-20931",} 4.3035456E7
# HELP vendor_jgroups_cluster_mfc_get_number_of_credit_requests_sent Number of credit requests sent
# TYPE vendor_jgroups_cluster_mfc_get_number_of_credit_requests_sent gauge
vendor_jgroups_cluster_mfc_get_number_of_credit_requests_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_num_timer_tasks Number of timer tasks queued up for execution
# TYPE vendor_jgroups_cluster_tcp_get_num_timer_tasks gauge
vendor_jgroups_cluster_tcp_get_num_timer_tasks{node="fercoli-mac-20931",} 7.0
# HELP vendor_cache_manager_default_cache_container_stats_remove_misses Cache container total number of cache removals where keys were not found
# TYPE vendor_cache_manager_default_cache_container_stats_remove_misses gauge
vendor_cache_manager_default_cache_container_stats_remove_misses{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_deliverable_messages Total number of deliverable messages in all receive windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_deliverable_messages gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_deliverable_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_fd_sock_get_num_suspected_members The number of currently suspected members
# TYPE vendor_jgroups_cluster_fd_sock_get_num_suspected_members gauge
vendor_jgroups_cluster_fd_sock_get_num_suspected_members{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_server_single_port_11222_transport_pending_tasks Returns the number of pending tasks.
# TYPE vendor_cache_manager_default_server_single_port_11222_transport_pending_tasks gauge
vendor_cache_manager_default_server_single_port_11222_transport_pending_tasks{node="fercoli-mac-20931",} 0.0
# HELP vendor_cpu_systemLoadAverage Displays the system load average for the last minute. The system load average is the sum of the number of runnable entities queued to the available processors and the number of runnable entities running on the available processors averaged over a period of time. The way in which the load average is calculated is operating system specific but is typically a damped time-dependent average. If the load average is not available, a negative value is displayed. This attribute is designed to provide a hint about the system load and may be queried frequently. The load average may be unavailable on some platform where it is expensive to implement this method.
# TYPE vendor_cpu_systemLoadAverage gauge
vendor_cpu_systemLoadAverage 0.89013671875
# HELP vendor_cache_manager_default_cluster_container_stats_memory_max The maximum amount of memory that JVMs across the cluster will attempt to utilise in bytes
# TYPE vendor_cache_manager_default_cluster_container_stats_memory_max gauge
vendor_cache_manager_default_cluster_container_stats_memory_max{node="fercoli-mac-20931",} 5.36870912E8
# HELP vendor_jgroups_cluster_unicast3_get_age_out_cache_size 
# TYPE vendor_jgroups_cluster_unicast3_get_age_out_cache_size gauge
vendor_jgroups_cluster_unicast3_get_age_out_cache_size{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_hit_ratio Cache container total percentage hit/(hit+miss) ratio for this cache
# TYPE vendor_cache_manager_default_cache_container_stats_hit_ratio gauge
vendor_cache_manager_default_cache_container_stats_hit_ratio{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_num_current_rows Prints the number of rows currently allocated in the matrix. This value will not be lower than xmit_table_now_rows
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_num_current_rows gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_num_current_rows{node="fercoli-mac-20931",} 50.0
# HELP vendor_jgroups_cluster_tcp_get_internal_thread_pool_size Current number of threads in the internal thread pool
# TYPE vendor_jgroups_cluster_tcp_get_internal_thread_pool_size gauge
vendor_jgroups_cluster_tcp_get_internal_thread_pool_size{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_messages_received 
# TYPE vendor_jgroups_cluster_unicast3_get_num_messages_received gauge
vendor_jgroups_cluster_unicast3_get_num_messages_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_server_single_port_11222_transport_number_of_global_connections Returns a count of active connections in the cluster. This operation will make remote calls to aggregate results, so latency might have an impact on the speed of calculation of this attribute.
# TYPE vendor_cache_manager_default_server_single_port_11222_transport_number_of_global_connections gauge
vendor_cache_manager_default_server_single_port_11222_transport_number_of_global_connections{node="fercoli-mac-20931",} 2.0
# HELP vendor_jgroups_cluster_stable_get_stable_received 
# TYPE vendor_jgroups_cluster_stable_get_stable_received gauge
vendor_jgroups_cluster_stable_get_stable_received{node="fercoli-mac-20931",} 39.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_undelivered_messages Total number of undelivered messages in all receive windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_undelivered_messages gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_undelivered_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_stable_get_stable_sent 
# TYPE vendor_jgroups_cluster_stable_get_stable_sent gauge
vendor_jgroups_cluster_stable_get_stable_sent{node="fercoli-mac-20931",} 39.0
# HELP vendor_jgroups_cluster_mfc_get_number_of_blockings Number of times flow control blocks sender
# TYPE vendor_jgroups_cluster_mfc_get_number_of_blockings gauge
vendor_jgroups_cluster_mfc_get_number_of_blockings{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_read_write_ratio Cache container read/writes ratio in all caches from this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_read_write_ratio gauge
vendor_cache_manager_default_cache_container_stats_read_write_ratio{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_become_server_queue_size_actual Actual size of the become_server_queue
# TYPE vendor_jgroups_cluster_nakack2_get_become_server_queue_size_actual gauge
vendor_jgroups_cluster_nakack2_get_become_server_queue_size_actual{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_stores Cache container total number of cache put operations
# TYPE vendor_cache_manager_default_cache_container_stats_stores gauge
vendor_cache_manager_default_cache_container_stats_stores{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_data_memory_used Amount in bytes of memory used in a given cache container for entries with eviction
# TYPE vendor_cache_manager_default_cache_container_stats_data_memory_used gauge
vendor_cache_manager_default_cache_container_stats_data_memory_used{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_num_purges Number of retransmit table purges
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_num_purges gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_num_purges{node="fercoli-mac-20931",} 1.0
# HELP vendor_jgroups_cluster_mfc_get_number_of_credit_responses_received Number of credit responses received
# TYPE vendor_jgroups_cluster_mfc_get_number_of_credit_responses_received gauge
vendor_jgroups_cluster_mfc_get_number_of_credit_responses_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_ufc_get_average_time_blocked Average time blocked (in ms) in flow control when trying to send a message
# TYPE vendor_jgroups_cluster_ufc_get_average_time_blocked gauge
vendor_jgroups_cluster_ufc_get_average_time_blocked{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_gms_get_number_of_views 
# TYPE vendor_jgroups_cluster_gms_get_number_of_views gauge
vendor_jgroups_cluster_gms_get_number_of_views{node="fercoli-mac-20931",} 1.0
# HELP vendor_cache_manager_default_cache_container_stats_average_read_time Cache container total average number of milliseconds for all read operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_read_time gauge
vendor_cache_manager_default_cache_container_stats_average_read_time{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cluster_container_stats_memory_total The total amount of memory in the JVMs across the cluster in bytes
# TYPE vendor_cache_manager_default_cluster_container_stats_memory_total gauge
vendor_cache_manager_default_cluster_container_stats_memory_total{node="fercoli-mac-20931",} 8.0740352E7
# HELP vendor_jgroups_cluster_tcp_get_thread_pool_size_active Current number of active threads in the thread pool
# TYPE vendor_jgroups_cluster_tcp_get_thread_pool_size_active gauge
vendor_jgroups_cluster_tcp_get_thread_pool_size_active{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_bundler_num_spins Number of spins before a real lock is acquired
# TYPE vendor_jgroups_cluster_tcp_bundler_num_spins gauge
vendor_jgroups_cluster_tcp_bundler_num_spins{node="fercoli-mac-20931",} 5.0
# HELP vendor_cache_manager_default_cache_container_health_free_memory_kb The amount of free memory (KB) in the host
# TYPE vendor_cache_manager_default_cache_container_health_free_memory_kb gauge
vendor_cache_manager_default_cache_container_health_free_memory_kb{node="fercoli-mac-20931",} 36712.0
# HELP vendor_jgroups_cluster_fd_sock_get_client_bind_port_actual The actual client_bind_port
# TYPE vendor_jgroups_cluster_fd_sock_get_client_bind_port_actual gauge
vendor_jgroups_cluster_fd_sock_get_client_bind_port_actual{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_gms_get_view_handler_size 
# TYPE vendor_jgroups_cluster_gms_get_view_handler_size gauge
vendor_jgroups_cluster_gms_get_view_handler_size{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_gms_get_num_members 
# TYPE vendor_jgroups_cluster_gms_get_num_members gauge
vendor_jgroups_cluster_gms_get_num_members{node="fercoli-mac-20931",} 1.0
# HELP vendor_jgroups_cluster_stable_get_stability_sent 
# TYPE vendor_jgroups_cluster_stable_get_stability_sent gauge
vendor_jgroups_cluster_stable_get_stability_sent{node="fercoli-mac-20931",} 39.0
# HELP vendor_jgroups_cluster_ufc_get_number_of_credit_requests_sent Number of credit requests sent
# TYPE vendor_jgroups_cluster_ufc_get_number_of_credit_requests_sent gauge
vendor_jgroups_cluster_ufc_get_number_of_credit_requests_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_remove_hits Cache container total number of cache removal hits
# TYPE vendor_cache_manager_default_cache_container_stats_remove_hits gauge
vendor_cache_manager_default_cache_container_stats_remove_hits{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_evictions Cache container total number of cache eviction operations
# TYPE vendor_cache_manager_default_cache_container_stats_evictions gauge
vendor_cache_manager_default_cache_container_stats_evictions{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_unacked_messages 
# TYPE vendor_jgroups_cluster_unicast3_get_num_unacked_messages gauge
vendor_jgroups_cluster_unicast3_get_num_unacked_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_size_of_all_messages Returns the number of bytes of all messages in all retransmit buffers. To compute the size, Message.getLength() is used
# TYPE vendor_jgroups_cluster_nakack2_get_size_of_all_messages gauge
vendor_jgroups_cluster_nakack2_get_size_of_all_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_fd_sock_get_num_suspect_events_generated Number of suspect event generated
# TYPE vendor_jgroups_cluster_fd_sock_get_num_suspect_events_generated gauge
vendor_jgroups_cluster_fd_sock_get_num_suspect_events_generated{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_thread_pool_size_largest Largest number of threads in the thread pool
# TYPE vendor_jgroups_cluster_tcp_get_thread_pool_size_largest gauge
vendor_jgroups_cluster_tcp_get_thread_pool_size_largest{node="fercoli-mac-20931",} 3.0
# HELP vendor_jgroups_cluster_unicast3_get_num_acks_received 
# TYPE vendor_jgroups_cluster_unicast3_get_num_acks_received gauge
vendor_jgroups_cluster_unicast3_get_num_acks_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_mfc_get_average_time_blocked Average time blocked (in ms) in flow control when trying to send a message
# TYPE vendor_jgroups_cluster_mfc_get_average_time_blocked gauge
vendor_jgroups_cluster_mfc_get_average_time_blocked{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cluster_container_stats_memory_available The maximum amount of free memory in bytes across the cluster JVMs
# TYPE vendor_cache_manager_default_cluster_container_stats_memory_available gauge
vendor_cache_manager_default_cluster_container_stats_memory_available{node="fercoli-mac-20931",} 3.7704896E7
# HELP vendor_jgroups_cluster_tcp_get_bundler_buffer_size 
# TYPE vendor_jgroups_cluster_tcp_get_bundler_buffer_size gauge
vendor_jgroups_cluster_tcp_get_bundler_buffer_size{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_average_read_time_nanos Cache container total average number of nanoseconds for all read operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_read_time_nanos gauge
vendor_cache_manager_default_cache_container_stats_average_read_time_nanos{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_acks_sent 
# TYPE vendor_jgroups_cluster_unicast3_get_num_acks_sent gauge
vendor_jgroups_cluster_unicast3_get_num_acks_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_average_write_time_nanos Cache container average number of nanoseconds for all write operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_write_time_nanos gauge
vendor_cache_manager_default_cache_container_stats_average_write_time_nanos{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_health_number_of_cpus Number of CPUs in the host
# TYPE vendor_cache_manager_default_cache_container_health_number_of_cpus gauge
vendor_cache_manager_default_cache_container_health_number_of_cpus{node="fercoli-mac-20931",} 12.0
# HELP vendor_jgroups_cluster_unicast3_get_num_connections Returns the total number of outgoing (send) and incoming (receive) connections
# TYPE vendor_jgroups_cluster_unicast3_get_num_connections gauge
vendor_jgroups_cluster_unicast3_get_num_connections{node="fercoli-mac-20931",} 0.0
# HELP vendor_cpu_availableProcessors Displays the number of processors available to the Java virtual machine. This value may change during a particular invocation of the virtual machine.
# TYPE vendor_cpu_availableProcessors gauge
vendor_cpu_availableProcessors 12.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_num_compactions Number of retransmit table compactions
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_num_compactions gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_num_compactions{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_read Returns the total number of bytes read by the server from clients which includes both protocol and user information.
# TYPE vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_read gauge
vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_read{node="fercoli-mac-20931",} 3084.0
# HELP vendor_cache_manager_default_cache_container_stats_current_number_of_entries_in_memory Cache container total number of entries currently in-memory for all caches in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_current_number_of_entries_in_memory gauge
vendor_cache_manager_default_cache_container_stats_current_number_of_entries_in_memory{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_time_since_start Number of seconds since cache started
# TYPE vendor_cache_manager_default_cache_container_stats_time_since_start gauge
vendor_cache_manager_default_cache_container_stats_time_since_start{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_size_of_all_messages_incl_headers Returns the number of bytes of all messages in all retransmit buffers. To compute the size, Message.size() is used
# TYPE vendor_jgroups_cluster_nakack2_get_size_of_all_messages_incl_headers gauge
vendor_jgroups_cluster_nakack2_get_size_of_all_messages_incl_headers{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_required_minimum_number_of_nodes Required minimum number of nodes to hold current cache data
# TYPE vendor_cache_manager_default_cache_container_stats_required_minimum_number_of_nodes gauge
vendor_cache_manager_default_cache_container_stats_required_minimum_number_of_nodes{node="fercoli-mac-20931",} -1.0
# HELP vendor_jgroups_cluster_ufc_get_number_of_credit_responses_sent Number of credit responses sent
# TYPE vendor_jgroups_cluster_ufc_get_number_of_credit_responses_sent gauge
vendor_jgroups_cluster_ufc_get_number_of_credit_responses_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_open_connections 
# TYPE vendor_jgroups_cluster_tcp_get_open_connections gauge
vendor_jgroups_cluster_tcp_get_open_connections{node="fercoli-mac-20931",} 0.0
# HELP vendor_cpu_processCpuTime Displays the CPU time used by the process on which the Java virtual machine is running in nanoseconds
# TYPE vendor_cpu_processCpuTime gauge
vendor_cpu_processCpuTime 5.7890516575E-314
# HELP vendor_cache_manager_default_cache_container_stats_average_remove_time Cache container total average number of milliseconds for all remove operation in this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_average_remove_time gauge
vendor_cache_manager_default_cache_container_stats_average_remove_time{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_num_resizes Number of retransmit table resizes
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_num_resizes gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_num_resizes{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_num_moves Number of moves in all (receive and send) windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_num_moves gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_num_moves{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cluster_container_stats_time_since_reset Number of seconds since the cluster-wide statistics were last reset
# TYPE vendor_cache_manager_default_cluster_container_stats_time_since_reset gauge
vendor_cache_manager_default_cluster_container_stats_time_since_reset{node="fercoli-mac-20931",} 46728.0
# HELP vendor_jgroups_cluster_stable_get_stability_received 
# TYPE vendor_jgroups_cluster_stable_get_stability_received gauge
vendor_jgroups_cluster_stable_get_stability_received{node="fercoli-mac-20931",} 39.0
# HELP vendor_jgroups_cluster_unicast3_get_num_messages_sent 
# TYPE vendor_jgroups_cluster_unicast3_get_num_messages_sent gauge
vendor_jgroups_cluster_unicast3_get_num_messages_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_current_seqno 
# TYPE vendor_jgroups_cluster_nakack2_get_current_seqno gauge
vendor_jgroups_cluster_nakack2_get_current_seqno{node="fercoli-mac-20931",} 8.0
# HELP vendor_jgroups_cluster_ufc_get_number_of_blockings Number of times flow control blocks sender
# TYPE vendor_jgroups_cluster_ufc_get_number_of_blockings gauge
vendor_jgroups_cluster_ufc_get_number_of_blockings{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_timer_threads Number of threads currently in the pool
# TYPE vendor_jgroups_cluster_tcp_get_timer_threads gauge
vendor_jgroups_cluster_tcp_get_timer_threads{node="fercoli-mac-20931",} 2.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_missing_messages Total number of missing messages in all receive windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_missing_messages gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_missing_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_different_cluster_messages Number of messages from members in a different cluster
# TYPE vendor_jgroups_cluster_tcp_get_different_cluster_messages gauge
vendor_jgroups_cluster_tcp_get_different_cluster_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_number_of_thread_dumps Number of thread dumps
# TYPE vendor_jgroups_cluster_tcp_get_number_of_thread_dumps gauge
vendor_jgroups_cluster_tcp_get_number_of_thread_dumps{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_stats_misses Cache container total number of cache attribute misses
# TYPE vendor_cache_manager_default_cache_container_stats_misses gauge
vendor_cache_manager_default_cache_container_stats_misses{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_cache_container_health_total_memory_kb The amount of total memory (KB) in the host
# TYPE vendor_cache_manager_default_cache_container_health_total_memory_kb gauge
vendor_cache_manager_default_cache_container_health_total_memory_kb{node="fercoli-mac-20931",} 78848.0
# HELP vendor_jgroups_cluster_ufc_get_number_of_credit_responses_received Number of credit responses received
# TYPE vendor_jgroups_cluster_ufc_get_number_of_credit_responses_received gauge
vendor_jgroups_cluster_ufc_get_number_of_credit_responses_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_missing_messages Total number of missing (= not received) messages in all retransmit buffers
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_missing_messages gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_missing_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_mfc_get_number_of_credit_requests_received Number of credit requests received
# TYPE vendor_jgroups_cluster_mfc_get_number_of_credit_requests_received gauge
vendor_jgroups_cluster_mfc_get_number_of_credit_requests_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_non_member_messages Number of messages from non-members
# TYPE vendor_jgroups_cluster_nakack2_get_non_member_messages gauge
vendor_jgroups_cluster_nakack2_get_non_member_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_send_connections Returns the number of outgoing (send) connections
# TYPE vendor_jgroups_cluster_unicast3_get_num_send_connections gauge
vendor_jgroups_cluster_unicast3_get_num_send_connections{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_different_version_messages Number of messages from members with a different JGroups version
# TYPE vendor_jgroups_cluster_tcp_get_different_version_messages gauge
vendor_jgroups_cluster_tcp_get_different_version_messages{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_merge3_get_views Number of cached ViewIds
# TYPE vendor_jgroups_cluster_merge3_get_views gauge
vendor_jgroups_cluster_merge3_get_views{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_server_single_port_11222_transport_number_of_local_connections Returns a count of active connections this server.
# TYPE vendor_cache_manager_default_server_single_port_11222_transport_number_of_local_connections gauge
vendor_cache_manager_default_server_single_port_11222_transport_number_of_local_connections{node="fercoli-mac-20931",} 2.0
# HELP vendor_cache_manager_default_cluster_size Size of the cluster in number of nodes
# TYPE vendor_cache_manager_default_cluster_size gauge
vendor_cache_manager_default_cluster_size{node="fercoli-mac-20931",} 1.0
# HELP vendor_jgroups_cluster_ufc_get_number_of_credit_requests_received Number of credit requests received
# TYPE vendor_jgroups_cluster_ufc_get_number_of_credit_requests_received gauge
vendor_jgroups_cluster_ufc_get_number_of_credit_requests_received{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_num_resizes Number of resizes in all (receive and send) windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_num_resizes gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_num_resizes{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_nakack2_get_xmit_table_undelivered_msgs Total number of undelivered messages in all retransmit buffers
# TYPE vendor_jgroups_cluster_nakack2_get_xmit_table_undelivered_msgs gauge
vendor_jgroups_cluster_nakack2_get_xmit_table_undelivered_msgs{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_tcp_get_internal_thread_pool_size_largest Largest number of threads in the internal thread pool
# TYPE vendor_jgroups_cluster_tcp_get_internal_thread_pool_size_largest gauge
vendor_jgroups_cluster_tcp_get_internal_thread_pool_size_largest{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_xmit_table_num_compactions Number of compactions in all (receive and send) windows
# TYPE vendor_jgroups_cluster_unicast3_get_xmit_table_num_compactions gauge
vendor_jgroups_cluster_unicast3_get_xmit_table_num_compactions{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_unicast3_get_num_xmits 
# TYPE vendor_jgroups_cluster_unicast3_get_num_xmits gauge
vendor_jgroups_cluster_unicast3_get_num_xmits{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_number_of_created_caches The total number of created caches, including the default cache.
# TYPE vendor_cache_manager_default_number_of_created_caches gauge
vendor_cache_manager_default_number_of_created_caches{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_mfc_get_number_of_credit_responses_sent Number of credit responses sent
# TYPE vendor_jgroups_cluster_mfc_get_number_of_credit_responses_sent gauge
vendor_jgroups_cluster_mfc_get_number_of_credit_responses_sent{node="fercoli-mac-20931",} 0.0
# HELP vendor_jgroups_cluster_stable_get_num_votes The number of votes for the current digest
# TYPE vendor_jgroups_cluster_stable_get_num_votes gauge
vendor_jgroups_cluster_stable_get_num_votes{node="fercoli-mac-20931",} 0.0
# HELP vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_written Returns the total number of bytes written by the server back to clients which includes both protocol and user information.
# TYPE vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_written gauge
vendor_cache_manager_default_server_single_port_11222_transport_total_bytes_written{node="fercoli-mac-20931",} 5312.0
# HELP vendor_cache_manager_default_cache_container_stats_number_of_entries Cache container total number of entries currently in all caches from this cache container
# TYPE vendor_cache_manager_default_cache_container_stats_number_of_entries gauge
vendor_cache_manager_default_cache_container_stats_number_of_entries{node="fercoli-mac-20931",} 0.0