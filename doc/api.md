Documentation for scan reporting API used by LanSweeper Agent

[[_TOC_]]

## Typical flow

| Local         | Cloud                 | Description                          |
| ------------- | --------------------- | ------------------------------------ |
| `AssetStatus` | `GetAssetStatus`      | Check if this asset is still enabled |
| `Config`      | `GetAssetGroupConfig` | Request updated client configuration |
| `Update`      | `GetUpdate`           | Request updated agent (buggy)        |
| `ScanData`    | `AddAssetScan`        | Submit scan result                   |

## Notes

### On `Content-Type` for requests to local scan servers

LSAgent requests specify `Content-Type: application/x-www-form-urlencoded` but send multipart MIME in the body.

I suspect this is unintentional, and likely a result of the vendor deciding the submitted JSON was too large, so they decided to gzip it, and then found URL encoded form data is unsuitable for submitting a binary file, and therefore switch to multipart MIME encoding — but forgot to update their header.

The Microsoft-HTTPAPI service responds with the correct `Content-Type: multipart/form-data; boundary=---`.

### AssetId and AgentKey

**AgentKey** is a shared secret used to submit reports to the relay server. This secret allows any client to submit scan results and, I suspect, the scan server to collect them.

A MIME field named **AgentKey** is included in requests sent to local scan servers but the value is set to the randomly generated AssetId. I suspect this is a bug.

**AgentKey** is set correctly in requests sent to the cloud relay.

### Updates

The client doesn't appear to know its own version if the `Version=` line is deleted from `LsAgent.ini` and will then report its version as 7.0.20.1 even though the installed version is 7.2.110.19.

This causes the client to request an update from the local scan server. The local scan server responds `OK` but doesn't provide an update.

The cloud relay server responds that the current version is 7.2.110.18 and then when the client requests the installer it says it doesn't exist and dumps a stack trace.

## Hello

Haven't captured this action while testing.

## AssetStatus

### Local

#### Request

```
POST /lsagent HTTP/1.1
Connection: Keep-Alive
Content-Type: application/x-www-form-urlencoded
Content-Length: 696
Host: lansweeper.example.org:9524
```

| Name            | Value                                |
| --------------- | ------------------------------------ |
| AgentKey        | b8447573-f51f-49d3-91be-ea5c310de84b |
| OperatingSystem | Linux                                |
| AssetId         | b8447573-f51f-49d3-91be-ea5c310de84b |
| Action          | AssetStatus                          |

#### Response

```
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: multipart/form-data; boundary=----------deadbeefdeadbeefdeadbeefdeadbeef
Server: Microsoft-HTTPAPI/2.0
Date: Fri, 04 Jun 2021 10:45:22 GMT
content-length: 252
```

| Name   | Value   |
| ------ | ------- |
| Status | Enabled |

### Cloud relay

#### Request

```
POST /EchoService.svc HTTP/1.1
Cache-Control: no-cache, max-age=0
SOAPAction: "urn:IEchoService/GetAssetStatus"
Accept-Encoding: gzip, deflate
Content-Type: text/xml; charset=utf-8
Content-Length: 239
Host: lsagentrelay.lansweeper.com

<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Body>
        <GetAssetStatus>
            <agentKey>b8447573-f51f-49d3-91be-ea5c310de84b</agentKey>
            <assetId>d0e3403c-2227-494f-9dcf-adf7d1eadd14</assetId>
        </GetAssetStatus>
    </s:Body>
</s:Envelope>
```

#### Response

```
HTTP/1.1 200 OK
Cache-Control: private
Content-Length: 195
Content-Type: text/xml; charset=utf-8
Vary: Accept-Encoding
Server: Microsoft-IIS/10.0
X-AspNet-Version: 4.0.30319
Request-Context: appId=cid-v1:06f0a4c9-9272-4566-a188-f07e7e7b52a5
Access-Control-Expose-Headers: Request-Context
X-Powered-By: ASP.NET
Date: Fri, 04 Jun 2021 13:01:40 GMT

<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Body>
        <GetAssetStatusResponse>
            <GetAssetStatusResult>Enabled</GetAssetStatusResult>
        </GetAssetStatusResponse>
    </s:Body>
</s:Envelope>
```

## Config

### Request

```
POST /lsagent HTTP/1.1
Connection: Keep-Alive
Content-Type: application/x-www-form-urlencoded
Content-Length: 691
Host: lansweeper.example.org:9524
```

| Name            | Value                                |
| --------------- | ------------------------------------ |
| AgentKey        | 06f0a4c9-9272-4566-a188-f07e7e7b52a5 |
| OperatingSystem | Linux                                |
| AssetId         | 06f0a4c9-9272-4566-a188-f07e7e7b52a5 |
| Action          | Config                               |

### Response

```
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: multipart/form-data; boundary=----------deadbeefdeadbeefdeadbeefdeadbeef
Server: Microsoft-HTTPAPI/2.0
Date: Fri, 04 Jun 2021 10:45:22 GMT
content-length: 21191
```

| Name   | Filename | Value                                         |
| ------ | -------- | --------------------------------------------- |
| Config | Config   | [see xml](doc/sample-data/lsagent.config.xml) |

## Update

### Request

| Name            | Value                                |
| --------------- | ------------------------------------ |
| AgentKey        | 00000000-0000-0000-0000-000000000000 |
| OperatingSystem | Linux                                |
| AssetId         | 67aa0432-166f-4f33-8de3-b316ba562f2c |
| ClientVersion   | 7.0.20.1                             |
| Action          | Update                               |

### Response

```
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Server: Microsoft-HTTPAPI/2.0
Date: Fri, 04 Jun 2021 10:45:27 GMT
content-length: 4

OK\r\n
```

**Note** Local server does not appear to attempt to _provide_ an update. Pretty sure it should respond with a field/file named `Installer`.

## ScanData

### Request

```
POST /lsagent HTTP/1.1
Connection: Keep-Alive
Content-Type: application/x-www-form-urlencoded
Content-Length: 39013
Host: lansweeper.example.org:9524
```

| Name            | Filename | Value                                |
| --------------- | -------- | ------------------------------------ |
| AgentKey        |          | b8447573-f51f-49d3-91be-ea5c310de84b |
| OperatingSystem |          | Linux                                |
| AssetId         |          | b8447573-f51f-49d3-91be-ea5c310de84b |
| Action          |          | ScanData                             |
| Scan            | Scan     | see xml                              |

### Response

```
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Server: Microsoft-HTTPAPI/2.0
Date: Fri, 04 Jun 2021 10:45:27 GMT
content-length: 4

OK\r\n
```

**Note** No `content-type` specified for this response.
