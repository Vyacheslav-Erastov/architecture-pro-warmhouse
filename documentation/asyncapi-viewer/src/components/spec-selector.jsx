import { Box, FormControl, InputLabel, MenuItem, Select, Typography } from '@mui/material'
import React from 'react'

export const SpecSelector = ({ specs, selectedSpec, onSpecChange }) => {
  return (
    <Box>
      <FormControl margin='dense'>
        <InputLabel>Спецификация</InputLabel>
        <Select
          value={selectedSpec?.file || ''}
          label="Спецификация"
          onChange={(e) => {
            const spec = specs.find(s => s.file === e.target.value)
            if (spec) onSpecChange(spec)
          }}
        >
          {specs.map(spec => (
            <MenuItem key={spec.file} value={spec.file}>{spec.name}</MenuItem>
          ))}
        </Select>
      </FormControl>

      <Box>
        {selectedSpec && (
          <Typography >Файл: {selectedSpec.file}</Typography>
        )}
      </Box>
    </Box>
  )
}
