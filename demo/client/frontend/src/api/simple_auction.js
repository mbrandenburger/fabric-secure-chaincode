/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

import api from "./api";

const invoke = "/cc/invoke";
const query = "/cc/query";

const auctionId = "SpectrumAuction2020";

export default {
  buildPayload(tx, args) {
    return {
      tx: tx,
      args: [JSON.stringify(args)]
    };
  },

  ///////////////////////////////////////////////
  // Helpers
  ///////////////////////////////////////////////

  getDefaultAuction() {
    return api.get("/clock_auction/getDefaultAuction");
  },

  ///////////////////////////////////////////////
  // Auctioneer actions
  ///////////////////////////////////////////////

  createAuction(args) {
    // just take the name of the auction and remove spaces
    args = {
      tx: "create",
      args: [auctionId]
    };

    return api.post(`${invoke}`, args);
  },

  getAuctionDetails() {
    return api.get(`/clock_auction/getAuctionDetails/` + auctionId);
  },

  getAuctionStatus() {
    let args = {
      tx: "status",
      args: [auctionId]
    };

    return api.post(`${query}`, args);
  },

  startNextRound(args) {
    args = {
      tx: "eval",
      args: [auctionId]
    };
    return api.post(`${invoke}`, args);
  },

  endRound(args) {
    args = {
      tx: "close",
      args: [auctionId]
    };

    return api.post(`${invoke}`, args);
  },

  ///////////////////////////////////////////////
  // Bidder actions
  ///////////////////////////////////////////////

  submitClockBid(args) {
    function getSum(total, num) {
      return total + Math.round(num.price * num.qty);
    }
    let bid = args.bids.reduce(getSum, 0);

    args = {
      tx: "submit",
      args: [auctionId, args.bidder, bid.toString()]
    };

    return api.post(`${invoke}`, args);
  },

  getRoundInfo(args) {
    return api.post(`${query}`, this.buildPayload("getRoundInfo", args));
  },

  getBidderRoundResults(args) {
    return api.post(
      `${query}`,
      this.buildPayload("getBidderRoundResults", args)
    );
  },

  getOwnerRoundResults(args) {
    return api.post(
      `${query}`,
      this.buildPayload("getOwnerRoundResults", args)
    );
  },

  submitAssignBid(args) {
    return api.post(`${invoke}`, this.buildPayload("submitAssignBid", args));
  },

  getAssignmentResults(args) {
    return api.post(
      `${query}`,
      this.buildPayload("getAssignmentResults", args)
    );
  }
};
