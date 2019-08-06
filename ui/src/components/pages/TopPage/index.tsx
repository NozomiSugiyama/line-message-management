import AddIcon from "@material-ui/icons/Add";
import React, { useContext } from "react";
import { RouteChildrenProps } from "react-router";
import Fab from "src/components/atoms/Fab";
import Page from "src/components/atoms/Page";
import Header from "src/components/molecules/Header";
import RouterHistoryContext from "src/contexts/RouterHistoryContext";

export type TopPageProps = React.ComponentProps<typeof Page> & RouteChildrenProps<{problemId: string}>;

export default (props: TopPageProps) => {
    const routerHistory = useContext(RouterHistoryContext);

    return (
        <Page
            {...props}
        >
            <Header
                appTitle="Top"
            />
            <div>Test</div>
            <Fab
                onClick={() => routerHistory.history.push("/top")}
            >
                <AddIcon />
            </Fab>
        </Page>
    );
};
