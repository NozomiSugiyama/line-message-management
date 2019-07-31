import AddIcon from "@material-ui/icons/Add";
import React, { useContext } from "react";
import { RouteChildrenProps } from "react-router";
import Fab from "src/components/atoms/Fab";
import Header from "src/components/molecules/Header";
import Host from "src/components/pages/TopPage/Host";
import RouterHistoryContext from "src/contexts/RouterHistoryContext";

export type WorkListPageProps = React.ComponentProps<typeof Host> & RouteChildrenProps<{problemId: string}>;

export default (props: WorkListPageProps) => {
    const routerHistory = useContext(RouterHistoryContext);

    return (
        <Host
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
        </Host>
    );
};
