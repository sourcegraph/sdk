package grafana_test

import (
	"encoding/json"
	"fmt"

	"github.com/grafov/autograf/grafana"
)

func ExampleNewBoard() {
	board := grafana.NewBoard("Sample dashboard title")
	row1 := board.AddRow("Sample row title")
	row1.AddGraph(grafana.NewGraph("Sample graph title"))
	data, _ := json.MarshalIndent(board, "", "  ")
	fmt.Printf("%s", data)
	// Output:
	// {
	//   "id": 1,
	//   "title": "Sample dashboard title",
	//   "originalTitle": "",
	//   "tags": null,
	//   "style": "dark",
	//   "timezone": "browser",
	//   "editable": true,
	//   "hideControls": false,
	//   "sharedCrosshair": false,
	//   "rows": [
	//     {
	//       "title": "Sample row title",
	//       "collapse": false,
	//       "editable": true,
	//       "height": "250px",
	//       "panels": [
	//         {
	//           "id": 1,
	//           "title": "Sample graph title",
	//           "type": "graph",
	//           "error": false,
	//           "isNew": true,
	//           "renderer": "flot",
	//           "aliasColors": null,
	//           "bars": false,
	//           "fill": 0,
	//           "grid": {
	//             "leftLogBase": null,
	//             "leftMax": null,
	//             "leftMin": null,
	//             "rightLogBase": null,
	//             "rightMax": null,
	//             "rightMin": null,
	//             "threshold1": null,
	//             "threshold1Color": "",
	//             "threshold2": null,
	//             "threshold2Color": "",
	//             "thresholdLine": false
	//           },
	//           "legend": {
	//             "avg": false,
	//             "current": false,
	//             "max": false,
	//             "min": false,
	//             "show": false,
	//             "total": false,
	//             "values": false
	//           },
	//           "lines": false,
	//           "linewidth": 0,
	//           "nullPointMode": "connected",
	//           "percentage": false,
	//           "pointradius": 5,
	//           "points": false,
	//           "seriesOverrides": null,
	//           "span": 12,
	//           "stack": false,
	//           "steppedLine": false,
	//           "timeFrom": null,
	//           "timeShift": null,
	//           "tooltip": {
	//             "shared": false,
	//             "value_type": ""
	//           },
	//           "x-axis": true,
	//           "y-axis": true,
	//           "y_formats": null
	//         }
	//       ]
	//     }
	//   ],
	//   "templating": {
	//     "list": null
	//   },
	//   "annotations": {
	//     "list": null
	//   },
	//   "schemiaVersion": 0,
	//   "version": 0,
	//   "links": null,
	//   "time": {
	//     "from": "",
	//     "to": ""
	//   },
	//   "timepicker": {
	//     "refresh_intervals": null,
	//     "time_options": null
	//   }
	// }

}