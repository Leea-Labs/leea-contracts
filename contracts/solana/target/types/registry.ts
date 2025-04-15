/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/registry.json`.
 */
export type Registry = {
  "address": "4JzjcmJSqWi2gpk8tZT93JHtyMimn3bubQMYs8FLAopM",
  "metadata": {
    "name": "registry",
    "version": "0.1.0",
    "spec": "0.1.0",
    "description": "Leea-Labs Agent Registry"
  },
  "instructions": [
    {
      "name": "registerAgent",
      "discriminator": [
        135,
        157,
        66,
        195,
        2,
        113,
        175,
        30
      ],
      "accounts": [
        {
          "name": "holder",
          "writable": true,
          "signer": true
        },
        {
          "name": "agentAccount",
          "writable": true,
          "pda": {
            "seeds": [
              {
                "kind": "const",
                "value": [
                  108,
                  101,
                  101,
                  97,
                  95,
                  97,
                  103,
                  101,
                  110,
                  116
                ]
              },
              {
                "kind": "account",
                "path": "holder"
              }
            ]
          }
        },
        {
          "name": "systemProgram",
          "address": "11111111111111111111111111111111"
        }
      ],
      "args": [
        {
          "name": "agentName",
          "type": "string"
        },
        {
          "name": "description",
          "type": "string"
        },
        {
          "name": "fee",
          "type": "u64"
        }
      ]
    },
    {
      "name": "registerAgentByAdmin",
      "discriminator": [
        33,
        145,
        96,
        28,
        167,
        153,
        151,
        193
      ],
      "accounts": [
        {
          "name": "admin",
          "writable": true,
          "signer": true,
          "address": "NAnMTmTdCYoszmR7cWsd1DodzfUCadnjXoHoHqwkPak"
        },
        {
          "name": "holder"
        },
        {
          "name": "agentAccount",
          "writable": true,
          "pda": {
            "seeds": [
              {
                "kind": "const",
                "value": [
                  108,
                  101,
                  101,
                  97,
                  95,
                  97,
                  103,
                  101,
                  110,
                  116
                ]
              },
              {
                "kind": "account",
                "path": "holder"
              }
            ]
          }
        },
        {
          "name": "systemProgram",
          "address": "11111111111111111111111111111111"
        }
      ],
      "args": [
        {
          "name": "agentName",
          "type": "string"
        },
        {
          "name": "description",
          "type": "string"
        },
        {
          "name": "fee",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateAgentScore",
      "discriminator": [
        87,
        172,
        57,
        159,
        88,
        130,
        103,
        191
      ],
      "accounts": [
        {
          "name": "admin",
          "writable": true,
          "signer": true,
          "address": "NAnMTmTdCYoszmR7cWsd1DodzfUCadnjXoHoHqwkPak"
        },
        {
          "name": "holder",
          "writable": true
        },
        {
          "name": "agentAccount",
          "writable": true,
          "pda": {
            "seeds": [
              {
                "kind": "const",
                "value": [
                  108,
                  101,
                  101,
                  97,
                  95,
                  97,
                  103,
                  101,
                  110,
                  116
                ]
              },
              {
                "kind": "account",
                "path": "holder"
              }
            ]
          }
        },
        {
          "name": "systemProgram",
          "address": "11111111111111111111111111111111"
        }
      ],
      "args": [
        {
          "name": "score",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "agentAccount",
      "discriminator": [
        241,
        119,
        69,
        140,
        233,
        9,
        112,
        50
      ]
    }
  ],
  "types": [
    {
      "name": "agentAccount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "holder",
            "type": "pubkey"
          },
          {
            "name": "agentName",
            "type": "string"
          },
          {
            "name": "description",
            "type": "string"
          },
          {
            "name": "fee",
            "type": "u64"
          },
          {
            "name": "score",
            "type": "u64"
          },
          {
            "name": "createdAt",
            "type": "i64"
          },
          {
            "name": "updatedAt",
            "type": "i64"
          }
        ]
      }
    }
  ]
};
