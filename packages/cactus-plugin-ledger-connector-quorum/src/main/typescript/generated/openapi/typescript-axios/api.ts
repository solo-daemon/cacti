/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Connector Quorum
 * Can perform basic tasks on a Quorum ledger
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface DeployContractSolidityBytecodeV1Request
 */
export interface DeployContractSolidityBytecodeV1Request {
    /**
     * The contract name for retrieve the contracts json on the keychain.
     * @type {string}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    contractName: string;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    web3SigningCredential: Web3SigningCredential;
    /**
     * See https://ethereum.stackexchange.com/a/47556 regarding the maximum length of the bytecode
     * @type {string}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    bytecode: string;
    /**
     * The keychainId for retrieve the contracts json.
     * @type {string}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    keychainId: string;
    /**
     * 
     * @type {number}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    gas?: number;
    /**
     * 
     * @type {string}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    gasPrice?: string;
    /**
     * The amount of milliseconds to wait for a transaction receipt with theaddress of the contract(which indicates successful deployment) beforegiving up and crashing.
     * @type {number}
     * @memberof DeployContractSolidityBytecodeV1Request
     */
    timeoutMs?: number;
}
/**
 * 
 * @export
 * @interface DeployContractSolidityBytecodeV1Response
 */
export interface DeployContractSolidityBytecodeV1Response {
    /**
     * 
     * @type {Web3TransactionReceipt}
     * @memberof DeployContractSolidityBytecodeV1Response
     */
    transactionReceipt: Web3TransactionReceipt;
}
/**
 * 
 * @export
 * @enum {string}
 */
export enum EthContractInvocationType {
    Send = 'SEND',
    Call = 'CALL'
}

/**
 * 
 * @export
 * @interface InvokeContractV1Request
 */
export interface InvokeContractV1Request {
    /**
     * 
     * @type {string}
     * @memberof InvokeContractV1Request
     */
    contractName: string;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof InvokeContractV1Request
     */
    signingCredential: Web3SigningCredential;
    /**
     * 
     * @type {EthContractInvocationType}
     * @memberof InvokeContractV1Request
     */
    invocationType: EthContractInvocationType;
    /**
     * The name of the contract method to invoke.
     * @type {string}
     * @memberof InvokeContractV1Request
     */
    methodName: string;
    /**
     * The list of arguments to pass in to the contract method being invoked.
     * @type {Array<any>}
     * @memberof InvokeContractV1Request
     */
    params: Array<any>;
    /**
     * The application binary interface of the solidity contract, optional parameter
     * @type {Array<any>}
     * @memberof InvokeContractV1Request
     */
    contractAbi?: Array<any>;
    /**
     * Address of the solidity contract, optional parameter
     * @type {string}
     * @memberof InvokeContractV1Request
     */
    contractAddress?: string;
    /**
     * 
     * @type {string | number}
     * @memberof InvokeContractV1Request
     */
    value?: string | number;
    /**
     * 
     * @type {string | number}
     * @memberof InvokeContractV1Request
     */
    gas?: string | number;
    /**
     * 
     * @type {string | number}
     * @memberof InvokeContractV1Request
     */
    gasPrice?: string | number;
    /**
     * 
     * @type {number}
     * @memberof InvokeContractV1Request
     */
    nonce?: number;
    /**
     * The amount of milliseconds to wait for a transaction receipt beforegiving up and crashing. Only has any effect if the invocation type is SEND
     * @type {number}
     * @memberof InvokeContractV1Request
     */
    timeoutMs?: number;
    /**
     * The keychainId for retrieve the contracts json.
     * @type {string}
     * @memberof InvokeContractV1Request
     */
    keychainId?: string;
}
/**
 * 
 * @export
 * @interface InvokeContractV1Response
 */
export interface InvokeContractV1Response {
    /**
     * 
     * @type {Web3TransactionReceipt}
     * @memberof InvokeContractV1Response
     */
    transactionReceipt?: Web3TransactionReceipt;
    /**
     * 
     * @type {any}
     * @memberof InvokeContractV1Response
     */
    callOutput?: any | null;
    /**
     * 
     * @type {boolean}
     * @memberof InvokeContractV1Response
     */
    success: boolean;
}
/**
 * 
 * @export
 * @interface QuorumTransactionConfig
 */
export interface QuorumTransactionConfig {
    [key: string]: object | any;

    /**
     * 
     * @type {string}
     * @memberof QuorumTransactionConfig
     */
    rawTransaction?: string;
    /**
     * 
     * @type {string | number}
     * @memberof QuorumTransactionConfig
     */
    from?: string | number;
    /**
     * 
     * @type {string}
     * @memberof QuorumTransactionConfig
     */
    to?: string;
    /**
     * 
     * @type {string | number}
     * @memberof QuorumTransactionConfig
     */
    value?: string | number;
    /**
     * 
     * @type {string | number}
     * @memberof QuorumTransactionConfig
     */
    gas?: string | number;
    /**
     * 
     * @type {string | number}
     * @memberof QuorumTransactionConfig
     */
    gasPrice?: string | number;
    /**
     * 
     * @type {number}
     * @memberof QuorumTransactionConfig
     */
    nonce?: number;
    /**
     * 
     * @type {string}
     * @memberof QuorumTransactionConfig
     */
    data?: string;
}
/**
 * 
 * @export
 * @interface RunTransactionRequest
 */
export interface RunTransactionRequest {
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof RunTransactionRequest
     */
    web3SigningCredential: Web3SigningCredential;
    /**
     * 
     * @type {QuorumTransactionConfig}
     * @memberof RunTransactionRequest
     */
    transactionConfig: QuorumTransactionConfig;
    /**
     * The amount of milliseconds to wait for a transaction receipt with thehash of the transaction(which indicates successful execution) beforegiving up and crashing.
     * @type {number}
     * @memberof RunTransactionRequest
     */
    timeoutMs?: number;
}
/**
 * 
 * @export
 * @interface RunTransactionResponse
 */
export interface RunTransactionResponse {
    /**
     * 
     * @type {Web3TransactionReceipt}
     * @memberof RunTransactionResponse
     */
    transactionReceipt: Web3TransactionReceipt;
}
/**
 * 
 * @export
 * @interface SolidityContractJsonArtifact
 */
export interface SolidityContractJsonArtifact {
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    contractName: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    metadata?: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    bytecode?: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    deployedBytecode?: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    sourceMap?: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    deployedSourceMap?: string;
    /**
     * 
     * @type {string}
     * @memberof SolidityContractJsonArtifact
     */
    sourcePath?: string;
    /**
     * 
     * @type {{ [key: string]: object; }}
     * @memberof SolidityContractJsonArtifact
     */
    compiler?: { [key: string]: object; };
    /**
     * 
     * @type {{ [key: string]: object; }}
     * @memberof SolidityContractJsonArtifact
     */
    functionHashes?: { [key: string]: object; };
    /**
     * 
     * @type {object}
     * @memberof SolidityContractJsonArtifact
     */
    gasEstimates?: object;
}
/**
 * @type Web3SigningCredential
 * @export
 */
export type Web3SigningCredential = Web3SigningCredentialCactusKeychainRef | Web3SigningCredentialGethKeychainPassword | Web3SigningCredentialNone | Web3SigningCredentialPrivateKeyHex;

/**
 * 
 * @export
 * @interface Web3SigningCredentialCactusKeychainRef
 */
export interface Web3SigningCredentialCactusKeychainRef {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    type: Web3SigningCredentialType;
    /**
     * The ethereum account (public key) that the credential  belongs to. Basically the username in the traditional  terminology of authentication.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    ethAccount: string;
    /**
     * The key to use when looking up the the keychain entry holding the secret pointed to by the  keychainEntryKey parameter.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    keychainEntryKey: string;
    /**
     * The keychain ID to use when looking up the the keychain plugin instance that will be used to retrieve the secret pointed to by the keychainEntryKey parameter.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    keychainId: string;
}
/**
 * 
 * @export
 * @interface Web3SigningCredentialGethKeychainPassword
 */
export interface Web3SigningCredentialGethKeychainPassword {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialGethKeychainPassword
     */
    type: Web3SigningCredentialType;
    /**
     * The ethereum account (public key) that the credential  belongs to. Basically the username in the traditional terminology of authentication.
     * @type {string}
     * @memberof Web3SigningCredentialGethKeychainPassword
     */
    ethAccount: string;
    /**
     * A geth keychain unlock password.
     * @type {string}
     * @memberof Web3SigningCredentialGethKeychainPassword
     */
    secret: string;
}
/**
 * Using this denotes that there is no signing required because the transaction is pre-signed.
 * @export
 * @interface Web3SigningCredentialNone
 */
export interface Web3SigningCredentialNone {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialNone
     */
    type: Web3SigningCredentialType;
}
/**
 * 
 * @export
 * @interface Web3SigningCredentialPrivateKeyHex
 */
export interface Web3SigningCredentialPrivateKeyHex {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    type: Web3SigningCredentialType;
    /**
     * The ethereum account (public key) that the credential belongs to. Basically the username in the traditional terminology of authentication.
     * @type {string}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    ethAccount: string;
    /**
     * The HEX encoded private key of an eth account.
     * @type {string}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    secret: string;
}
/**
 * 
 * @export
 * @enum {string}
 */
export enum Web3SigningCredentialType {
    CactusKeychainRef = 'CACTUS_KEYCHAIN_REF',
    GethKeychainPassword = 'GETH_KEYCHAIN_PASSWORD',
    PrivateKeyHex = 'PRIVATE_KEY_HEX',
    None = 'NONE'
}

/**
 * 
 * @export
 * @interface Web3TransactionReceipt
 */
export interface Web3TransactionReceipt {
    [key: string]: object | any;

    /**
     * 
     * @type {boolean}
     * @memberof Web3TransactionReceipt
     */
    status: boolean;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionReceipt
     */
    transactionHash: string;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionReceipt
     */
    transactionIndex: number;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionReceipt
     */
    blockHash: string;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionReceipt
     */
    blockNumber: number;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionReceipt
     */
    gasUsed: number;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionReceipt
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionReceipt
     */
    from: string;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionReceipt
     */
    to: string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Deploys the bytecode of a Solidity contract.
         * @param {DeployContractSolidityBytecodeV1Request} [deployContractSolidityBytecodeV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumDeployContractSolidityBytecode: async (deployContractSolidityBytecodeV1Request?: DeployContractSolidityBytecodeV1Request, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-quorum/deploy-contract-solidity-bytecode`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(deployContractSolidityBytecodeV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Invokes a contract on a besu ledger
         * @param {InvokeContractV1Request} [invokeContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumInvokeContract: async (invokeContractV1Request?: InvokeContractV1Request, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-quorum/invoke-contract`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(invokeContractV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Executes a transaction on a quorum ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumRunTransaction: async (runTransactionRequest?: RunTransactionRequest, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-quorum/run-transaction`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(runTransactionRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusExporterMetricsV1: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-quorum/get-prometheus-exporter-metrics`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Deploys the bytecode of a Solidity contract.
         * @param {DeployContractSolidityBytecodeV1Request} [deployContractSolidityBytecodeV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request?: DeployContractSolidityBytecodeV1Request, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<DeployContractSolidityBytecodeV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Invokes a contract on a besu ledger
         * @param {InvokeContractV1Request} [invokeContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiV1QuorumInvokeContract(invokeContractV1Request?: InvokeContractV1Request, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<InvokeContractV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiV1QuorumInvokeContract(invokeContractV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Executes a transaction on a quorum ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiV1QuorumRunTransaction(runTransactionRequest?: RunTransactionRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RunTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiV1QuorumRunTransaction(runTransactionRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPrometheusExporterMetricsV1(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPrometheusExporterMetricsV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Deploys the bytecode of a Solidity contract.
         * @param {DeployContractSolidityBytecodeV1Request} [deployContractSolidityBytecodeV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request?: DeployContractSolidityBytecodeV1Request, options?: any): AxiosPromise<DeployContractSolidityBytecodeV1Response> {
            return localVarFp.apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Invokes a contract on a besu ledger
         * @param {InvokeContractV1Request} [invokeContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumInvokeContract(invokeContractV1Request?: InvokeContractV1Request, options?: any): AxiosPromise<InvokeContractV1Response> {
            return localVarFp.apiV1QuorumInvokeContract(invokeContractV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Executes a transaction on a quorum ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiV1QuorumRunTransaction(runTransactionRequest?: RunTransactionRequest, options?: any): AxiosPromise<RunTransactionResponse> {
            return localVarFp.apiV1QuorumRunTransaction(runTransactionRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusExporterMetricsV1(options?: any): AxiosPromise<string> {
            return localVarFp.getPrometheusExporterMetricsV1(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Deploys the bytecode of a Solidity contract.
     * @param {DeployContractSolidityBytecodeV1Request} [deployContractSolidityBytecodeV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request?: DeployContractSolidityBytecodeV1Request, options?: any) {
        return DefaultApiFp(this.configuration).apiV1QuorumDeployContractSolidityBytecode(deployContractSolidityBytecodeV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Invokes a contract on a besu ledger
     * @param {InvokeContractV1Request} [invokeContractV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public apiV1QuorumInvokeContract(invokeContractV1Request?: InvokeContractV1Request, options?: any) {
        return DefaultApiFp(this.configuration).apiV1QuorumInvokeContract(invokeContractV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Executes a transaction on a quorum ledger
     * @param {RunTransactionRequest} [runTransactionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public apiV1QuorumRunTransaction(runTransactionRequest?: RunTransactionRequest, options?: any) {
        return DefaultApiFp(this.configuration).apiV1QuorumRunTransaction(runTransactionRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the Prometheus Metrics
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPrometheusExporterMetricsV1(options?: any) {
        return DefaultApiFp(this.configuration).getPrometheusExporterMetricsV1(options).then((request) => request(this.axios, this.basePath));
    }
}


