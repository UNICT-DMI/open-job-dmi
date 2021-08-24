$( document ).ready(function() {
	$( "#azienda" ).blur(checkForm);
	$( "#email" ).blur(checkForm);
	$( "#ruolo" ).blur(checkForm);
	$( "#salario" ).blur(checkForm);
	$( "#descrizione" ).blur(checkForm);
	$( "#competenze" ).blur(checkForm);
	$( "#sede" ).blur(checkForm);
	$( "#benefits" ).blur(checkForm);
	$( "#full-time" ).click(checkForm);
	$( "#part-time" ).click(checkForm);

	$( '#submit' ).click(function(){
		if($( '#submit-success' ).hasClass('d-none') == false)
			$( '#submit-success' ).addClass('d-none');
		
		if($( '#submit-error' ).hasClass('d-none') == false)
			$( '#submit-error' ).addClass('d-none');
		
		if($( '#email-error' ).hasClass('d-none') == false)
			$( '#email-error' ).addClass('d-none');

		dataRequest = createJsonForm();

		if(validateEmail(dataRequest['email']) == false){
			$( '#email-error' ).removeClass('d-none');
			return;
		}

		recaptcha_key = document.getElementById('recaptcha_key').value;

		grecaptcha.ready(function() {
			grecaptcha.execute(recaptcha_key, {action: 'submit'}).then(function(token) {
				dataRequest['recaptcha_token'] = token;

				$.ajax({
					url: '/api/offer',
					type: 'POST',
					dataType: 'json',
					data: JSON.stringify(dataRequest),
					error : function(xhr, textStatus, errorThrown) {
						$( '#submit-error' ).removeClass('d-none');
					},
					success:function(response){
						$( '#submit-success' ).removeClass('d-none');
					}
				});
			});
		  });
	});

	function createJsonForm(){
		var offerta = {}
		offerta['azienda'] = document.getElementById('azienda').value;
		offerta['ruolo'] = document.getElementById('ruolo').value;
		offerta['salario'] = document.getElementById('salario').value;
		offerta['sede'] = document.getElementById('sede').value;
		offerta['email'] = document.getElementById('email').value;
		offerta['descrizione'] = document.getElementById('descrizione').value;
		offerta['competenze'] = document.getElementById('competenze').value;
		offerta['benefits'] = document.getElementById('benefits').value;
		offerta['fulltime'] = document.getElementById('full-time').checked;
		offerta['parttime'] = document.getElementById('part-time').checked;

		return offerta;
	}

	function checkForm(){
		let submitDisabled = document.getElementById('submit').disabled;
		let offerta = createJsonForm();
		let error = false;
		
		for (let key in offerta){
			if(key == 'fulltime' || key == 'parttime' || key == 'benefits')
				continue;
			if(offerta[key] == ''){
				error = true;
				break;
			}
		}

		if(offerta['fulltime'] == false && offerta['parttime'] == false)
			error = true;
		
		if(error && submitDisabled == '')
			document.getElementById('submit').disabled = 'disabled';

		if(!error && submitDisabled != '')
			document.getElementById('submit').disabled = '';
		
	}

	function validateEmail(email) {
		const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
		return re.test(email);
	  }

});