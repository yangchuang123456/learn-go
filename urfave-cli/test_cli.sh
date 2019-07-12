#!/usr/bin/env bash

set -v

go build main.go

echo "success result"

./main --node_name default_b0 --data_dir /home/qydev/tmp/dipperin_apps/default_b0 --node_type 3 --p2p_listener :10003 --m_port 9202 --pprofport 20002 --http_port 10004 --ws_port 10005 --ipc_path /home/qydev/tmp/dipperin_apps/default_b0/dipperin.ipc --soft_wallet_pwd 123 --is_start_mine 1 --pprof true --debug_mode 2 --log_level info --is_upload_node_data 1 --upload_url http://127.0.0.1:8887/api/dipperin_nodes

echo "error result"
echo ""

./main --node_name default_b0 --data_dir /home/qydev/tmp/dipperin_apps/default_b0 --node_type 3 --p2p_listener :10003 --http_port 10004 --ws_port 10005 --ipc_path /home/qydev/tmp/dipperin_apps/default_b0/dipperin.ipc --soft_wallet_pwd 123 --is_start_mine 1 --pprof true --debug_mode 2 --log_level info --is_upload_node_data 1 --upload_url http://127.0.0.1:8887/api/dipperin_nodes --m_port 9202 --pprofport 20002
