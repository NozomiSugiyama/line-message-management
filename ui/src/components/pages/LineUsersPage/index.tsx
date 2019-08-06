import { Dialog, DialogContent, DialogTitle, List, ListItem, ListItemText, DialogActions, Button } from "@material-ui/core";
import React, { useEffect, useState } from "react";
import { RouteChildrenProps } from "react-router";
import { read } from "src/api/line_users";
import { LineUser, LineUsers } from "src/api/type";
import Page from "src/components/atoms/Page";
import Header from "src/components/molecules/Header";
import LabeledTypography from "src/components/molecules/LabeledTypography";
import styled from "styled-components";
import ChatWidget from "src/components/molecules/ChatWidget";
import send_message from "src/api/line_users/send_message";
import FlexibleSpace from "src/components/atoms/FlexibleSpace";
import send_test_messages from "src/api/line_users/send_test_messages";

export type LineUsersPageProps = React.ComponentProps<typeof Page> & RouteChildrenProps<{problemId: string}>;

export default (props: LineUsersPageProps) => {
    const [lineUsers, setLineUsers] = useState<LineUsers>([]);
    const [selectedLineUser,  selectLineUser] = useState<LineUser | null>(null);

    useEffect(
        () => {
            (async () => {
                setLineUsers(await read());
            })();
        },
        []
    );

    return (
        <Page
            {...props}
        >
            <Header
                appTitle="Line Users"
            />
            <List>
                {lineUsers.map(lineUser => (
                    <ListItem
                        key={lineUser.id}
                        button
                        onClick={() => {
                            selectLineUser(lineUser);
                        }}
                    >
                        <ListItemText primary={lineUser.user.name}/>
                    </ListItem>
                ))}
            </List>
            <Dialog
                open={!!selectedLineUser}
                onClose={() => selectLineUser(null)}
            >
                <DialogTitle>{selectedLineUser && selectedLineUser.user.name}</DialogTitle>
                <StyledDialogContent>
                    <div>
                        <LabeledTypography label="ID">{selectedLineUser && selectedLineUser.user.id}</LabeledTypography>
                        <LabeledTypography label="Email">{selectedLineUser && selectedLineUser.user.email}</LabeledTypography>
                        <LabeledTypography label="Name">{selectedLineUser && selectedLineUser.user.name}</LabeledTypography>
                    </div>
                    <div>
                        <LabeledTypography label="Linked Account">{selectedLineUser && selectedLineUser.linked_account}</LabeledTypography>
                        <LabeledTypography label="Line ID">{selectedLineUser && selectedLineUser.line_id}</LabeledTypography>
                    </div>
                </StyledDialogContent>
                <DialogContent>
                    <ChatWidget
                        onSubmit={(message) => {
                            if (!selectedLineUser) {
                                return;
                            }
                            send_message({ message, id: selectedLineUser.id });
                        }}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => selectLineUser(null)}>閉じる</Button>
                    <FlexibleSpace/>
                    <Button
                        onClick={() => {
                            if (!selectedLineUser) {
                                return;
                            }
                            send_test_messages({ id: selectedLineUser.id });
                        }}
                    >
                        Send Test Messages
                    </Button>
                </DialogActions>
            </Dialog>
        </Page>
    );
};

const StyledDialogContent = styled(DialogContent)`
    display: flex;
    > :first-child {
        margin-right: 2rem;
    }
    > :last-child {
        flex-grow: 1;
    }
`;
