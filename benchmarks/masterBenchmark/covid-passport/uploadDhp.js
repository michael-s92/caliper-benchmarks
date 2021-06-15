/* eslint-disable no-undef */
'use strict';

const seeds = require('./seeds-3x30.json');
const Utils = require('./utils');
const fetch = require('node-fetch');

class uploadDhp {

    static get() {

        (async () => {
            try{
            const response = await fetch("https://storage.googleapis.com/milan-thesis-21/covid-passport/seeds-3x30.json");
            const json = await response.json();

            console.log(json.url);
            console.log(json.explanation);

            }catch(err){
                console.log(err.response.body);
                return;
            }


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

        })();

      
	}
}

module.exports = uploadDhp;
