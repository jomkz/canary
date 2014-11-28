/**
* mail.js
*
* @description :: Service to send email.
*/

var nodemailer = require('nodemailer');
var transporter = nodemailer.createTransport({
    service: sails.config.email.service,
    auth: {
        user: sails.config.email.username,
        pass: sails.config.email.password
    }
});

module.exports = {

    sendIPUpdateEmail: function (address) {
        transporter.sendMail({
            from: sails.config.email.defaultFrom,
            to: sails.config.email.ipUpdateRecipients,
            subject: '[CANARY] IP Address Change Detected',
            text: 'An IP address change has been detected. The new IP address is ' + address.ipv4
        });
    }

};
