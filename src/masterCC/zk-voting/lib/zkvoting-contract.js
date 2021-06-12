/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');
const Election = require('./election');
/*
const Article = require('./article');
const { Author, Editor, Reviewer } = require('./users');
const ReviewingProcess = require('./reviewing_process');
const { ArticleSubmittedEvent, DoReviewEvent, ReviewDoneEvent } = require('./chaincode_events');
const Helper = require('./helper');
const sha512 = require('js-sha512');

const seeds = require('./seeds.json');

const authorTitleIndexName = "author~title";
const authorTitleReviewingIndexName = "author~title~reviewing";
*/
class ZKVotingContract extends Contract {

    async init(ctx) {
        console.info('ZKVoting Contract Initialized');
    }

    async doNothing(ctx) {
        console.info("DoNothing Transaction Invoked");
    }

    async initLedger(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async configureElection(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async closeElection(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async getCandidates(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async vote(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async getResults(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async hasVoted(ctx) {
        console.info("InitLedger Transaction Invoked");

    }

    async voteOK(ctx) {
        console.info("InitLedger Transaction Invoked");

    }
}

module.exports = ZKVotingContract
