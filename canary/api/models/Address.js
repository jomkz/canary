/**
* Address.js
*
* @description :: This model represents an IP address.
*/

module.exports = {

    attributes: {
        ipv4: {
            type: 'string',
            size: 15,
            unique: true
        }
    }
};
