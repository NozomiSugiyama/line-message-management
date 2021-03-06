import React, { Fragment } from "react";

// Make the minimum configuration.
export interface ErrorBoundaryState {
    error?: any;
    info?: any;
    hasError: boolean;
}

export default class ErrorBoundary extends React.Component<{ children?: React.ReactNode }, ErrorBoundaryState> {
    state: ErrorBoundaryState = { hasError: false };

    resetPage = () => {
        // Reset localStorage state
        localStorage.clear();
        // reload page
        location.reload();
    }

    componentDidCatch(error :any, info: any) {
        // Display fallback UI
        this.setState({ error, info, hasError: true });
        // You can also log the error to an error reporting service
        console.error(error, info);
    }

    render() {
        if (this.state.hasError) {
            // You can render any custom fallback UI
            return (
                <div
                    style={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                        flexDirection: "column",
                        margin: "5rem 0"
                    }}
                >
                    <h1>An unexpected error has occurred!</h1>
                    <button
                        onClick={this.resetPage}
                        style={{
                            display: "inline-block",
                            padding: "0.5em 1em",
                            textDecoration: "none",
                            borderRadius: "4px",
                            boxShadow: "rgba(0, 0, 0, 0.29) 0px 2px 9px 1px",
                            fontSize: "2rem",
                            margin: "2rem",
                            backgroundColor: "#fafbfd",
                            letterSpacing: "2px"
                        }}
                    >
                        RELOAD
                    </button>
                    <hr/>
                    <h2>Debug message for developer</h2>
                    <br/>
                    <pre style={{ fontSize: "1.5rem" }}>{this.state.error && (this.state.error.message || this.state.error.toString())}</pre>
                    <pre>
                        {React.isValidElement(this.state.info) ? this.state.info
                       :                                         Object.entries(this.state.info).map(([key, value]: [string, any]) => (
                            <div key={key}>
                                <div>{key}</div>
                                <div>
                                    {typeof value === "string"   ? value.split("\n").map(x => <Fragment key={x}>{x}<br/></Fragment>)
                                   : React.isValidElement(value) ? value
                                   :                               JSON.stringify(value)
                                    }
                                </div>
                            </div>
                        ))
                        }
                    </pre>
                </div>
            );
        }
        return this.props.children;
    }
}
