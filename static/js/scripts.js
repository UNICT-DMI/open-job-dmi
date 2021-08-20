$( document ).ready(function() {
	$( "#azienda" ).blur(checkForm);
	$( "#email" ).blur(checkForm);
	$( "#ruolo" ).blur(checkForm);
	$( "#descrizione" ).blur(checkForm);
	$( "#competenze" ).blur(checkForm);
	$( "#benefits" ).blur(checkForm);
	$( "#full-time" ).click(checkForm);
	$( "#part-time" ).click(checkForm);

	function createJsonForm(){
		var offerta = {}
		offerta['azienda'] = document.getElementById('azienda').value;
		offerta['email'] = document.getElementById('ruolo').value;
		offerta['descrizione'] = document.getElementById('descrizione').value;
		offerta['competenze'] = document.getElementById('competenze').value;
		offerta['benefits'] = document.getElementById('competenze').value;
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

});