# Demo to showcase dapr actor reminders.

- Install dapr, redis
- Install component definitions in comp.yaml
- Install okteto
- Deploy the manifest in deployment.yaml
- Run okteto up, and inside execute 
```
cd /app ; go run ./cmd/sub/
```
- Separate console, run okteto up -f oketo-pub.yaml
```
cd /app ; go run ./cmd/pub/main.go
```

This will execute 'EnterLane' and start a reminder to trigger in 10 seconds:

```
root@sub-okteto-57d5dfc5fb-xwjv9:/app# cd /app ; go run ./cmd/sub/
dapr client initializing for: 127.0.0.1:50001
2024/11/14 23:35:59 Car pl1 enter lane
2024/11/14 23:35:59 Reminder reg OK
2024/11/14 23:36:09 receive reminder = VehiLOST state=CAR3 duetime=10s period=
```

# Reminder lost if 'sub' crashes

- show logs of daprd in 'sub'
```
k logs -f sub-okteto-57d5dfc5fb-xwjv9 -c daprd
```

- Kill and re-execute 'pub'
```
root@pub-okteto-6cd45997f-f9bgb:/app# cd /app ; go run ./cmd/pub/main.go 
dapr client initializing for: 127.0.0.1:50001
2024/11/14 23:36:49 Get ID
```

- Before 10 seconds, kill 'sub'
```
root@sub-okteto-57d5dfc5fb-xwjv9:/app# cd /app ; go run ./cmd/sub/
dapr client initializing for: 127.0.0.1:50001
2024/11/14 23:39:08 Car pl1 enter lane
2024/11/14 23:39:08 Reminder reg OK
^Csignal: interrupt
root@sub-okteto-57d5dfc5fb-xwjv9:/app# 

```
- Dapr logs will  show:
```
 time="2024-11-14T23:39:08.297465385Z" level=debug msg="Starting to read reminders for actor type testActor2Type (migrate=false), with metadata id 00000000-0000-0000-0000-000000000000 and 0 partitions" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:08.297573844Z" level=debug msg="Read reminders from actors||testActor2Type without partition" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:08.297594511Z" level=debug msg="Finished reading reminders for actor type testActor2Type (migrate=false), with metadata id 00000000-0000-0000-0000-000000000000 and no partitions: total of 0 reminders" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:14.476486413Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:18.01813529Z" level=debug msg="Executing reminder for actor testActor2Type||CAR3||VehiLOST" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor type=log ver=1.14.4
time="2024-11-14T23:39:18.024640992Z" level=error msg="Error executing reminder for actor testActor2Type||CAR3||VehiLOST: Put \"http://127.0.0.1:8080/actors/testActor2Type/CAR3/method/remind/VehiLOST\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor type=log ver=1.14.4
time="2024-11-14T23:39:18.024734617Z" level=error msg="Error invoking reminder on actor testActor2Type||CAR3: Put \"http://127.0.0.1:8080/actors/testActor2Type/CAR3/method/remind/VehiLOST\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor type=log ver=1.14.4
time="2024-11-14T23:39:18.03808019Z" level=debug msg="Found reminder with key: testActor2Type||CAR3||VehiLOST. Deleting reminder" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:18.044418516Z" level=debug msg="Starting to read reminders for actor type testActor2Type (migrate=false), with metadata id 00000000-0000-0000-0000-000000000000 and 0 partitions" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:18.044544266Z" level=debug msg="Read reminders from actors||testActor2Type without partition" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:18.045994816Z" level=debug msg="Finished reading reminders for actor type testActor2Type (migrate=false), with metadata id 00000000-0000-0000-0000-000000000000 and no partitions: total of 1 reminders" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:19.475185468Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:24.476147327Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:29.47454459Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:29.474611923Z" level=warning msg="Actor health check failed 4 times, marking unhealthy" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:29.474772924Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:29.80315541Z" level=debug msg="Disconnecting from placement service by the unhealthy app" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:29.812069083Z" level=debug msg="Halting actor 'testActor2Type||CAR3'" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor type=log ver=1.14.4
time="2024-11-14T23:39:29.812616044Z" level=error msg="Failed to deactivate all actors: failed to deactivate actor 'testActor2Type||CAR3': Delete \"http://127.0.0.1:8080/actors/testActor2Type/CAR3\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:29.976371556Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:30.47633847Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
```

- restart 'sub' app. Note reminder is not received:
```
root@sub-okteto-57d5dfc5fb-xwjv9:/app# cd /app ; go run ./cmd/sub/
dapr client initializing for: 127.0.0.1:50001
```

- Dapr logs:
```
time="2024-11-14T23:39:37.476332941Z" level=error msg="Error performing request: Get \"http://127.0.0.1:8080/healthz\": dial tcp 127.0.0.1:8080: connect: connection refused" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:37.976269814Z" level=info msg="Actor health check succeeded, marking healthy" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=actorshealth type=log ver=1.14.4
time="2024-11-14T23:39:38.333009328Z" level=debug msg="try to connect to placement service: dns:///dapr-placement-server.dapr-system.svc.cluster.local:50005" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:38.343160716Z" level=debug msg="Established connection to placement service at dns:///dapr-placement-server.dapr-system.svc.cluster.local:50005" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:40.626865868Z" level=debug msg="Placement order received: lock" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:40.630257262Z" level=debug msg="Placement order received: update" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:40.630428763Z" level=info msg="Placement tables updated, version: 9" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:40.630443179Z" level=debug msg="Placement order received: unlock" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actors.placement type=log ver=1.14.4
time="2024-11-14T23:39:40.634513785Z" level=debug msg="Starting to read reminders for actor type testActor2Type (migrate=true), with metadata id 00000000-0000-0000-0000-000000000000 and 0 partitions" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:40.63545079Z" level=debug msg="Read reminders from actors||testActor2Type without partition" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:40.635542415Z" level=debug msg="Finished reading reminders for actor type testActor2Type (migrate=true), with metadata id 00000000-0000-0000-0000-000000000000 and no partitions: total of 0 reminders" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
time="2024-11-14T23:39:40.635559249Z" level=debug msg="Loaded 0 reminders for actor type testActor2Type" app_id=sub instance=sub-okteto-57d5dfc5fb-xwjv9 scope=dapr.runtime.actor.reminders type=log ver=1.14.4
```

