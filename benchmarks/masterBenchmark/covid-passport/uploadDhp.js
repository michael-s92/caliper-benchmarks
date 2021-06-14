/* eslint-disable no-undef */
'use strict';

const seeds = require('./seeds-3x30.json');
const Utils = require('./utils');

class uploadDhp {

    static get() {


        let dhps = seeds.validDHPs;

        let d = new Date();
        let randomAccessKey;
        do{
            randomAccessKey = Utils.getRandomInt(dhps.length) + d.getSeconds();
        } while(dhps[randomAccessKey] === undefined);

        let dhp = dhps[randomAccessKey];


        // PurgeExpiredDhps()
	    let args = {
                chaincodeFunction: 'uploadDhp',
                chaincodeArguments: [JSON.stringify(dhp)]
            };

	    return args;

	}
}

module.exports = uploadDhp;
