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
    },
    palette: {
        primary: {
            contrastText: "#fff",
            dark: "#c56200",
            light: "#ffc246",
            main: MAIN_COLOR,
        },
    }
});
