import React, { useEffect, useState } from 'react'
import './App.css'
import { Box, Typography } from "@mui/material"
import { SpecSelector, SpecViewer } from './components'

function App() {
  const [specs, setSpecs] = useState([])
  const [selectedSpec, setSelectedSpec] = useState(null)
  const [specContent, setSpecContent] = useState({ content: null, error: null })

  useEffect(() => {
    fetch("config.json")
      .then(response => {
        if (!response.ok) {
          throw new Error('Failed to load specification')
        }
        return response.json()
      })
      .then(content => {
        if (!content["specs"].length) {
          setSpecContent({ error: `Error loading specification\n\nFile config.json is invalid`, content: null })
        }
        setSpecs(content["specs"])
        setSelectedSpec(content["specs"][0])
      })
      .catch(error => {
        console.error('Error loading spec:', error)
        setSpecContent({ error: `Error loading specification\n\nCould not load: config.json`, content: null })
      })
  }, [])

  useEffect(() => {
    if (selectedSpec) {
      fetch(selectedSpec.file)
        .then(response => {
          if (!response.ok) {
            throw new Error('Failed to load specification')
          }
          return response.text()
        })
        .then(content => {
          setSpecContent({ error: null, content: content })
        })
        .catch(error => {
          console.error('Error loading spec:', error)
          setSpecContent({ error: `Error loading specification\n\nCould not load: ${selectedSpec.file}`, content: null })
        })
    }
  }, [selectedSpec])

  return (
    <Box sx={{ display: "flex", justifyContent: "center", flexDirection: "column", gap: 1, alignItems: "center", width: "100%", pt: 2 }}>
      <Box sx={{ textAlign: "center" }}>
        <Typography variant='h2'>AsyncAPI Specification Viewer</Typography>
        {specs.length !== 0 && (
          <SpecSelector
            specs={specs}
            selectedSpec={selectedSpec}
            onSpecChange={setSelectedSpec}
          />
        )}
      </Box>

      <Box>
        {specContent.error && (
          <Typography color='error'>{specContent.error}</Typography>
        )}
        {specContent.content && (
          <SpecViewer
            spec={specContent.content}
            specName={selectedSpec?.name}
          />
        )}
      </Box>
    </Box>
  )
}

export default App
