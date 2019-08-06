import { createMuiTheme } from "@material-ui/core/styles";

const MAIN_COLOR = "#27365d";

export default createMuiTheme({
    overrides: {
        MuiDrawer: {
            paper: {
                backgroundColor: "#fafbfd"
            },
            paperAnchorDockedLeft: {
                borderRight: "none"
            }
        },
        MuiDialog: {
            paper: {
                border: 0,
                borderRadius: 8,
                color: "#333",
            },
        },
        MuiListItem: {
            root: {
                borderRadius: 4
            }
        }
    },
    palette: {
        primary: {
            contrastText: "#fff",
            main: MAIN_COLOR,
        },
    }
});
