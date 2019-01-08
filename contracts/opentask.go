package contracts

const OpenTaskABI = `
[
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "accept",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "reject",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "missionId",
				"type": "string"
			}
		],
		"name": "publish",
		"outputs": [],
		"payable": true,
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "isOwner",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "x",
				"type": "bytes32"
			}
		],
		"name": "bytes32ToString",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			},
			{
				"name": "arbitrationId",
				"type": "string"
			}
		],
		"name": "confirm",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			},
			{
				"name": "missionId",
				"type": "string"
			},
			{
				"name": "data",
				"type": "string"
			}
		],
		"name": "solve",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newCurrency",
				"type": "string"
			}
		],
		"name": "updateCurrency",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "initialCurrency",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "missionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "reward",
				"type": "uint256"
			}
		],
		"name": "Publish",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "missionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "data",
				"type": "string"
			}
		],
		"name": "Solve",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "Accept",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "Reject",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "arbitrationId",
				"type": "string"
			}
		],
		"name": "Confirm",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	}
]`