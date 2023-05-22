/**
 * Hyperledger Cactus Plugin - Odap Hermes
 *
 * Implementation for Odap and Hermes
 *
 * The version of the OpenAPI document: 0.0.1
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
 * @param sessionID 
 * @param clientIdentityPubkey 
 * @param serverIdentityPubkey 
 * @param lockEvidenceClaim 
 * @param lockEvidenceExpiration 
 * @param hashCommenceAckRequest 
 * @param signature 
 * @param messageType 
 * @param sequenceNumber 
 * @param lockEvidenceFormat 
 * @param clientTransferNumber 
 * @param messageHash 
 */

data class LockEvidenceV1Request (

    @Json(name = "sessionID")
    val sessionID: kotlin.String,

    @Json(name = "clientIdentityPubkey")
    val clientIdentityPubkey: kotlin.String,

    @Json(name = "serverIdentityPubkey")
    val serverIdentityPubkey: kotlin.String,

    @Json(name = "lockEvidenceClaim")
    val lockEvidenceClaim: kotlin.String,

    @Json(name = "lockEvidenceExpiration")
    val lockEvidenceExpiration: kotlin.String,

    @Json(name = "hashCommenceAckRequest")
    val hashCommenceAckRequest: kotlin.String,

    @Json(name = "signature")
    val signature: kotlin.String,

    @Json(name = "messageType")
    val messageType: kotlin.String,

    @Json(name = "sequenceNumber")
    val sequenceNumber: java.math.BigDecimal,

    @Json(name = "lockEvidenceFormat")
    val lockEvidenceFormat: kotlin.Any? = null,

    @Json(name = "clientTransferNumber")
    val clientTransferNumber: kotlin.Int? = null,

    @Json(name = "messageHash")
    val messageHash: kotlin.String? = null

)
