// Helper object to interact with backend
var pulser = {
    // Subscribes a user
    subscribe: function (email) {
        return $.post('/subscribe', {email: email});
    },

    // Unsubscribes a user
    unsubscribe: function (email) {
        return $.post('/unsubscribe', {email: email});
    }
};

(function() {
    $(document).ready(function () {
        var $subscribeBtn = $('#subscribe');
        var $unsubscribeBtn = $('#unsubscribe');
        var $emailInput = $('#email');
        var $message = $('#message');
        var $toggleAction = $('.toggle-action');
        var $form = $('#form');

        // Validates email
        var validate = function (email) {
            if (!email) {
                $message.trigger('$message', {
                    statusCode: 500,
                    message: 'Email is required'
                });
                return false;
            }
            return true;
        };

        // When subscribe/unsubscribe events are successful, triggers message update
        var successHandler = function (data, status, r) {
            $message.trigger('$message', {
                statusCode: r.status,
                message: data.Message
            });
            $emailInput.val('');
            $form.hide();
        };

        // When subscribe/unsubscribe events go wrong, triggers message update
        var errorHandler = function (r) {
            $message.trigger('$message', {
                statusCode: r.status,
                message: r.responseText // simple error string is returned
            });
        };

        // Subscribe button handler
        $subscribeBtn.on('click', function () {
            var email = $emailInput.val();
            if (!validate(email)) {
                return;
            }
            pulser.subscribe(email).then(successHandler, errorHandler);
        });


        // Unsubscribe handler
        $unsubscribeBtn.on('click', function () {
            var email = $emailInput.val();
            if (!validate(email)) {
                return;
            }
            pulser.unsubscribe(email).then(successHandler, errorHandler);
        });

        $toggleAction.on('click', function () {
            if ($subscribeBtn.is(':visible')) {
                $subscribeBtn.hide();
                $unsubscribeBtn.show();
                $toggleAction.html('Or if you would like to subscribe, you can <a href="javascript:;">do so here</a>.');
            } else {
                $unsubscribeBtn.hide();
                $subscribeBtn.show();
                $toggleAction.html('Or if you\'re already subscribed, you can <a href="javascript:;">unsubscribe here</a>.');
            }
        });

        // Handle message updates from other handlers. Listens when `$message` event is emitted
        // and uses data in passed in payload to render information.
        // payload : {
        //  	message: string
        //  	statusCode: int
        // }
        $message.on('$message', function (e, payload) {
            if (payload.statusCode == 200) {
                $(this).removeClass('fail').addClass('success');
            } else {
                $(this).removeClass('success').addClass('fail');
            }
            $(this).text(payload.message);
        });
    });
})();
