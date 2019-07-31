import { createContext } from "react";

export type DrawerValue = {
    toggle: () => void,
    hide: () => void,
    show: () => void,
};

// It is declared by React Component
// To make the compilation successful, temporary values ​​are included
export default createContext<DrawerValue>({
    toggle: () => undefined,
    hide: () => undefined,
    show: () => undefined,
});
