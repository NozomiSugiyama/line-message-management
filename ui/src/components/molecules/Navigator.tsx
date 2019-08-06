import Divider from "@material-ui/core/Divider";
import Typography from "@material-ui/core/Typography";
import React, { useContext } from "react";
import Link from "src/components/atoms/Link";
import styled from "styled-components";
import { List, ListItem, ListItemText } from "@material-ui/core";
import RouterHistoryContext from "src/contexts/RouterHistoryContext";

export type NavigatorProps = React.ComponentProps<typeof Host>;

export default (props: NavigatorProps) => {
    const routerHistory = useContext(RouterHistoryContext);
    return (
        <Host {...props}>
            <Title variant="h2">
                <Link to="/">
                    LINE配信管理
                </Link>
            </Title>
            <Divider/>
            <List>
                <Link to="/line-users">
                    <ListItem
                        selected={/^\/line-users/.test(routerHistory.location.pathname)}
                        button
                    >
                        <ListItemText primary="Line Users" />
                    </ListItem>
                </Link>
                <Link to="/settings">
                    <ListItem
                        selected={/^\/settings/.test(routerHistory.location.pathname)}
                        button
                    >
                        <ListItemText primary="Settings" />
                    </ListItem>
                </Link>
            </List>
        </Host>
    );
};

const Host = styled.div`
    display: flex;
    flex-direction: column;
    min-height: calc(100vh - 2rem);
    max-height: calc(100vh - 2rem);
    width: 15rem;
    overflow: auto;
    margin: 1rem;
    box-sizing: border-box;
`;

const Title = styled(Typography)`
    && {
        font-size: 2rem;
        padding-top: 1.5rem;
        padding-bottom: 1rem;
        text-align: center;
        letter-spacing: .1rem;
    }
`;
