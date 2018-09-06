package telerik

import (
	"bytes"
)

type CustomTemplate struct {
	Id       []byte `jsObject:"-"`
	Template []byte `jsObject:"-"`
	Footer   []byte `jsObject:"-"`
	NoData   []byte `jsObject:"-"`
}

func (el *CustomTemplate) GetElementTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`kendo.template($('#`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`Template').html())`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`<script id='`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`Template' type='text/x-kendo-tmpl'>`))
	buffer.Write(el.Template)
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetElementFooterTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`kendo.template($('#`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`FooterTemplate').html())`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetFooterTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`<script id='`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`FooterTemplate' type='text/x-kendo-tmpl'>`))
	buffer.Write(el.Footer)
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetElementNoDataTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`kendo.template($('#`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`NoDataTemplate').html())`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetNoDataTemplate() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`<script id='`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`NoDataTemplate' type='text/x-kendo-tmpl'>`))
	buffer.Write(el.NoData)
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetFooterTemplateButton(buttonLabel, buttonCss []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`<button class='k-button k-primary `))
	buffer.Write(buttonCss)
	buffer.Write([]byte(`' onclick='`))
	buffer.Write(el.Id)
	buffer.Write([]byte(`CustomButtonClick()'>`))
	buffer.Write(buttonLabel)
	buffer.Write([]byte(`</button>`))

	return buffer.Bytes()
}
func (el *CustomTemplate) GetFooterTemplateButtonOpenDialogJavaScript() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(`function `))
	buffer.Write(el.Id)
	buffer.Write([]byte(`CustomButtonClick(){
      containerHostExposedPortsAddNewPortDialogWindowRef.kendoDialog({
        title: 'Expose port from container',
        content: kendo.template($('#containerCreateTemplateExposedPortsAddNewPort').html()),
        visible: false,
        modal: true,
        close: function(){

        },
        open: function(){
          $('#ExposedPortsNumber').kendoNumericTextBox({ format: '#' });
          containerConfigurationExposedPortsNumberRef = $('#ExposedPortsNumber').data('kendoNumericTextBox');

          $('#ExposedPortsProtocol').kendoDropDownList();

          setTimeout( function(){ containerConfigurationExposedPortsNumberRef.focus(); }, 1000)
        },
        actions: [
          {
            text: 'Close'
          },
          {
            text: 'Add',
            action: function(e){
              `))
	buffer.Write(el.Id)
	buffer.Write([]byte(`CustomButtonOnActionAndClose();
              return false;
            }
          },
          {
            text: 'Add and close',
            action: function(e){
              `))
	buffer.Write(el.Id)
	buffer.Write([]byte(`CustomButtonOnActionAndClose();
            },
            primary: true
          }
        ]
      });
      containerHostExposedPortsAddNewPortDialogWindowRef.data('kendoDialog').open();
    }
    function `))
	buffer.Write(el.Id)
	buffer.Write([]byte(`CustomButtonOnActionAndClose(){
      let imageName = containerConfigurationImageNameRef.text();

      // Procura por um nome de container
      if( imageName == '' ) {
        $('#dialog').kendoDialog({
          modal: true,
          visible: true,
          title: 'Configuration error',
          content: 'Please, select an image name first.',
          width: 400,
          actions: [
            { text: 'OK', primary: true, action: function(){ containerConfigurationImageNameRef.open(); } }
          ]
        });
        return;
      }

      // Procura por uma porta válida
      if( containerConfigurationExposedPortsNumberRef.value() == null ){
        $('#dialog').kendoDialog({
          modal: true,
          visible: true,
          title: 'Configuration error',
          content: 'Please, select a valid port number.',
          width: 400,
          actions: [
            { text: 'OK', primary: true, action: function(){ setTimeout( function(){ containerConfigurationExposedPortsNumberRef.focus(); }, 1000); } }
          ]
        });
        return;
      }

      let dataSource = containerHostExposedPortsRef.dataSource;
      containerHostExposedPortsItemsCount += 1;
      containerHostExposedPortsItemsIdToAdd = containerHostExposedPortsItemsCount;

      let dataInDataSource = dataSource.data();
      let pass = true;

      for (var i in dataInDataSource) {
        if( isNaN( i ) ){
          break
        }
        i = parseInt(i);
        if (dataInDataSource[i].Value === $('#ExposedPortsNumber').val() + '/' + $('#ExposedPortsProtocol').val()) {
          containerHostExposedPortsItemsIdToAdd = dataInDataSource[i].Id;
          pass = false;
          break;
        }
      }

      if( pass === true ) {
        dataSource.add({
          Id: containerHostExposedPortsItemsCount,
          Value: $('#ExposedPortsNumber').val() + '/' + $('#ExposedPortsProtocol').val(),
          ImageName: containerConfigurationImageNameRef.text()
        });
      }

      dataSource.one('requestEnd', function(args) {
        if (args.type !== 'create') {
          return;
        }

        dataSource.one('sync', function() {
          containerHostExposedPortsRef.value(containerHostExposedPortsRef.value().concat([containerHostExposedPortsItemsIdToAdd]));
        });

        $('#ExposedPortsNumber').data('kendoNumericTextBox').value('');
      });

      dataSource.sync();
    }`))

	return buffer.Bytes()
}
