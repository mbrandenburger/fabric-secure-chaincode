/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#pragma once

#include <stdbool.h>
#include <stdint.h>
#include <string>

typedef struct auction
{
    std::string name;
    bool is_open;
} auction_t;

typedef struct bid
{
    std::string bidder_name;
    int value;
} bid_t;

typedef struct status
{
    bool is_open;
    int num_bids;
} status_t;

typedef struct status_msg
{
    int rc;
    std::string message;
} status_msg_t;

int unmarshal_auction(auction_t* auction, const char* json_bytes, uint32_t json_len);
int unmarshal_bid(bid_t* bids, const char* json_bytes, uint32_t json_len);
std::string marshal_auction(auction_t* auction);
std::string marshal_bid(bid_t* bid);
std::string marshal_status(std::string state, int clockRound, bool roundActive);
std::string marshal_status_msg(status_msg_t* status_msg);
std::string marshal_response(status_msg_t* status_msg, std::string payload);
