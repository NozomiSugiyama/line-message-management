import Typography, { TypographyProps } from "@material-ui/core/Typography";
import React from "react";
import styled from "styled-components";

export interface LabeledTypography extends TypographyProps {
    label: string;
}

export default ({ label, ...props }: LabeledTypography) => (
    <Host>
        <span>{label}</span>
        <Typography {...props}/>
    </Host>
);

const Host = styled.div`
    display: flex;
    flex-direction: column;
    > :first-child {
        font-size: .5rem;
    }
    > :last-child {
        margin-top: -4px;
        margin-bottom: 4px;
        overflow: hidden;
        text-overflow: ellipsis;
        max-width: 12rem;
    }
`;
