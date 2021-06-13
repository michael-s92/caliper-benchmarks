/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');
const Election = require('./election');
const Admins = require('./admins');

const seeds = require('./seeds.json');
const sha512 = require('js-sha512');

const Helper = require('./helper');
/*
const { ArticleSubmittedEvent, DoReviewEvent, ReviewDoneEvent } = require('./chaincode_events');
const Helper = require('./helper');

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

        // save admins
        for (const adm of seeds.admins) {
            let hashedKey = sha512(adm.key);
            let objAdmin = new Admins(adm.id, hashedKey);
            await ctx.stub.putState(adm.id, Buffer.from(JSON.stringify(objAdmin)));
        }

        // save elections
        for (const e of seeds.initElections) {

            let obj = new Election(e.id, e.candidates);
            await ctx.stub.putState(e.id, Buffer.from(JSON.stringify(obj)));
        }

    }

    async configureElection(ctx, electionId, candidates, adminId, adminKey) {

        Helper.throwErrorIfStringIsEmpty(electionId);
        Helper.throwErrorIfStringIsEmpty(candidates);
        Helper.throwErrorIfStringIsEmpty(adminId);
        Helper.throwErrorIfStringIsEmpty(adminKey);

    }

    async closeElection(ctx, electionId, adminId, adminKey) {


    }

    async getCandidates(ctx, electionId) {


    }

    async vote(ctx, electionId, voterId, voteIndex) {


    }

    async getResults(ctx, electionId) {


    }

    async hasVoted(ctx, electionId, voterId) {


    }

    async voteOK(ctx, electionId, voterId) {

        if (electionId.length <= 0) {
            throw new Error("electionId must be non-empty string");
        }

        if (voterId.length <= 0) {
            throw new Error("voterId must be non-empty string");
        }


    }
}

module.exports = ZKVotingContract
