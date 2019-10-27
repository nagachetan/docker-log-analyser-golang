# docker-log-analyser-golang
This project is to tabulate the big set of logs coming out of docker containers or k8s pods. It will help to debug/analyse the logs

1. Desired log level can be input so that those log levels only will be filtered out and displayed in tabulated form.
2. Fonts of each column can be set with desired color
3. Currently 6 fields are added in the program this can be extended by adding more fields in the structure and extending the display of the table.
4. if it needs to be in html format it can be enabled with an option RenderHTML and it will generate the result into the html file.

