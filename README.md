# containerfs
a cluster filesystem for the containers

# Concepts

a volume = a metadata table + multiple block groups

# Key Problem: Ino Generation

auto increase on the client side, and fetch the initial ino from the metadata table on mount

or allocate on the metadata table side

# Replication

raft or paxos for the metadata table

append-only write for blockgroups


