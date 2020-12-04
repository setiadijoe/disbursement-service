# Disbursement Service
### Technical Testing Flip

There are three endpoint in this app. Here they are:
1. Get List Of Disbursement
2. Get Disbursement From Thirdparty
3. Update Disbursement From Thirdparty

| No. | Path | Method | Body |Description |
|-----|------|--------|------|------------|
|1.|"/disburse"|"GET" | |get list of the disbursement that has recorded|
|2.|"/disburse"|"POST"| "account_number" :string |get disbursement from third party|
||||"bank_code" :string ||
||||"amount" :float||
||||"remark" :string||
|3.|"/disburse/:id"|"PUT"||update a disbursement from third party|