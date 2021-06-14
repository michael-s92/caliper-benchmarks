/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');
const sha512 = require('js-sha512');

//const seeds = require('./seeds.json');


class CovidPassportContract extends Contract {

    async init(ctx) {
        console.info('CovidPassport Contract Initialized');
    }

    async doNothing(ctx) {
        console.log('=========== DoNothing Function');
    }

    async initLedger(ctx) {
        console.log('=========== START: initLedger Transaction');


        console.log('=========== END  : initLedger Transaction');
    }

    async UploadDhp(ctx) {

    }

    async VerifyResult(ctx) {

    }

    async PurgeExpiredDhps(ctx) {

    }

}

module.exports = CovidPassportContract;
