## axelard query gov param

Query the parameters (voting|tallying|deposit) of the governance process

### Synopsis

Query the all the parameters for the governance process.

Example:
$ <appd> query gov param voting
$ <appd> query gov param tallying
$ <appd> query gov param deposit

```
axelard query gov param [param-type] [flags]
```

### Options

```
      --height int    Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help          help for param
      --node string   <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID (default "axelar")
      --home string         directory for config and data (default "$HOME/.axelar")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string       Output format (text|json) (default "text")
      --trace               print out full stack trace on errors
```

### SEE ALSO

- [axelard query gov](axelard_query_gov.md)	 - Querying commands for the governance module
