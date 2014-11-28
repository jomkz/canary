/**
 * HomeController
 *
 * @description :: Controller to handle the default view for the application.
 */

module.exports = {

    index: function (req, res) {
        Address.find(function (err, addresses) {
            return res.view({ error: err, addresses: addresses});
        });
    }

};
