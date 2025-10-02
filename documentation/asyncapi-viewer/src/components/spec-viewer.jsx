import React from 'react'
import AsyncApiComponent from '@asyncapi/react-component'
import "@asyncapi/react-component/styles/default.css";
import { Box } from '@mui/material';

export const SpecViewer = ({ spec, specName }) => {

    return (
        <Box >
            <div className="asyncapi-container">
                <AsyncApiComponent
                    schema={spec}
                    config={{
                        show: {
                            sidebar: true,
                        },
                        expand: {
                            messageExamples: false,
                        }
                    }}
                />
            </div>
        </Box>
    )
}