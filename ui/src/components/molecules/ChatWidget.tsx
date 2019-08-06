import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import Paper from "@material-ui/core/Paper";
import SendIcon from "@material-ui/icons/Send";
import React from "react";
import styled from "styled-components";

const StyledPaper = styled(Paper)`
    min-width: 20rem;
    display: flex;
    align-items: center;
    width: 400;
`;

const StyledInputBase = styled(InputBase)`
    margin-left: 1rem;
    flex-grow: 1;
`;

export default ({ onSubmit }: { onSubmit: (chatMessage: string) => void }) => {
    return (
        <form
            onSubmit={async (e: React.FormEvent<HTMLFormElement>) => {
                e.preventDefault();

                const chatMessage = (e.target as any).elements["chat-message"].value;
                onSubmit(chatMessage);
            }}
        >
            <StyledPaper>
                <StyledInputBase
                    placeholder="Input Chat"
                    id="chat-message"
                    required
                />
                <IconButton
                    component="button"
                    type="submit"
                >
                    <SendIcon />
                </IconButton>
            </StyledPaper>
        </form>
    );
};
