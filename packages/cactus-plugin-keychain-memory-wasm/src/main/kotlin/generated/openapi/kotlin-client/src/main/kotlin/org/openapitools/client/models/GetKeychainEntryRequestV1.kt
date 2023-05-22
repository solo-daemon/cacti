/**
 * Hyperledger Cactus Plugin - Keychain Memory WASM
 *
 * Contains/describes the Hyperledger Cactus Keychain Memory WASM plugin which is designed to help with testing and development and is implemented in Rust instead of Typescript.
 *
 * The version of the OpenAPI document: 1.0.0-rc.1
 * 
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json

/**
 * 
 *
 * @param key The key for the entry to get from the keychain.
 */

data class GetKeychainEntryRequestV1 (

    /* The key for the entry to get from the keychain. */
    @Json(name = "key")
    val key: kotlin.String

)
