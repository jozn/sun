CREATE TABLE files (
    id bigint,
    data blob,
    data_thumb blob,
    file_type text,
    PRIMARY KEY (id)
) WITH
  comment=''
  AND read_repair_chance=0
  AND dclocal_read_repair_chance=0.1
  AND gc_grace_seconds=864000
  AND bloom_filter_fp_chance=0.01
  AND compaction={ 'class':'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy',
  'max_threshold':'32',
  'min_threshold':'4' }
  AND compression= {'enabled':'false'}
  AND caching={ 'keys':'ALL',
  'rows_per_partition':'NONE' }
  AND default_time_to_live=0
  AND id='b564b550-f7ef-11e7-a35b-894f7b0f4a8a'
  AND min_index_interval=128
  AND max_index_interval=2048
  AND memtable_flush_period_in_ms=0
  AND speculative_retry='99PERCENTILE';
