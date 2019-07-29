import Fab from "@material-ui/core/Fab";
import React from "react";
import styled from "styled-components";

export default (props: React.ComponentProps<typeof FixedFab>) => <FixedFab {...props}/>;

const FixedFab = styled(Fab)`
    && {
        position: fixed;
        right: 0;
        bottom: 0;
        margin: 2rem;
    }
`;
