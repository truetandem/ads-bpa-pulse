describe("A set of tests for subscribing and unsubscribing", function(){

	it("Checks if a new subscription can be created", function(done){
		var email = new Date().getTime() + '@local.dev';

		pulser.subscribe(email).then(function(data, s, r){
			expect(r.status).toEqual(200);
			done();
		});
	});

	it("Checks if a new subscription can be created and then removed", function(done){
		var email = new Date().getTime() + '@local.dev';

		pulser.subscribe(email).then(function(data, s, r){
			expect(r.status).toEqual(200);

			pulser.unsubscribe(email).then(function(data, s, r2){
				expect(r2.status).toEqual(200);
				done();
			});
		});
	});
});
