curl -k -D- https://lansweeper.example.org:9524/lsagent \
-F AgentKey=640f05b6-9ef4-45f8-b376-ab189ed48082 \
-F OperatingSystem=Linux \
-F AssetId=640f05b6-9ef4-45f8-b376-ab189ed48082 \
-F Action=ScanData \
-F "Scan=@out.gz;filename=Scan"
