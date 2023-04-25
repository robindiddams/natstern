# changefeed-consumers
checks the pending messages for all changefeed consumers

## install 
```bash
git clone github.com:robindiddams/changefeed-consumers.git
cd changefeed-consumers
go install .
```


## run
```bash
❯  changefeed-consumers
using credential <your creds file>
using url <url>
changefeed_consumer__AppointmentToUser pending: 0, ackPending: 0
changefeed_consumer_appointment pending: 60830, ackPending: 1
changefeed_consumer_authorization pending: 0, ackPending: 0
changefeed_consumer_authorization_service pending: 0, ackPending: 0
changefeed_consumer_billing_plan pending: 0, ackPending: 0
changefeed_consumer_billing_plan_add_on pending: 0, ackPending: 0
...
```

### filter out non-zero pending consumers
```bash
❯  changefeed-consumers nz
using credential  <your creds file>
using url <url>
changefeed_consumer_appointment pending: 60830, ackPending: 1
changefeed_consumer_entity_change pending: 74, ackPending: 1
changefeed_consumer_inspection pending: 241, ackPending: 1
changefeed_consumer_labor_matrix pending: 4, ackPending: 1
changefeed_consumer_message_thread_metadata pending: 21, ackPending: 1
changefeed_consumer_notification pending: 1782, ackPending: 1
changefeed_consumer_order pending: 17, ackPending: 1
changefeed_consumer_password_reset_token pending: 52, ackPending: 1
changefeed_consumer_service pending: 1211, ackPending: 1
changefeed_consumer_subcontract pending: 359, ackPending: 1
changefeed_consumer_vcdb_vehicle pending: 77750, ackPending: 2
changefeed_consumer_vehicle pending: 4, ackPending: 1

```
