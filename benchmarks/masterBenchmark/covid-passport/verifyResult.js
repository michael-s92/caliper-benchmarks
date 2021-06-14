/* eslint-disable no-undef */
/* eslint-disable no-unused-vars */
/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

'use strict';

const seeds = require('./seeds-3x30.json');
const Utils = require('./utils');

class verifyResult {

    static get() {

        let dhps = seeds.validDHPs;

        let randomAccessKey;
        do{
            randomAccessKey = Utils.getRandomInt(dhps.length);
        } while(dhps[randomAccessKey] === undefined);

        let dhp = dhps[randomAccessKey];


        // verifyResult(final Context ctx, final String patient, final String method)
	    let args = {
                chaincodeFunction: 'verifyResult',
                chaincodeArguments: [dhp.data.patient, dhp.data.method]
            };

	    return args;

	}
}

module.exports = verifyResult;




