# containerfs
a cluster filesystem for the containers

# Concepts

volume, a filesystem instance

blockgroup,

inodegroup,

# Key Problem: Ino Generation

alloc on the server side

directory inodes have even numbers and regular files use odd numbers.

# Memory, Disk, and Server

Metadata (inodes, directry trees) are totally in memory and persisted onto the disk. 

A datanode is associated with one disk, which make things simpler. And one data storage server may run multiple datanodes. 

# Replication

raft or paxos for the inodegroup

append-only write for blockgroups


