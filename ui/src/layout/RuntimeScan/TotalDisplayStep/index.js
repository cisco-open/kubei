import React from 'react';
import VulnerabilitiesSummaryDisplay, { CisBenchmarkLevelsDisplay } from 'components/VulnerabilitiesSummaryDisplay';
import StepDisplay, { StepDisplayTitle } from '../StepDisplay';

import './total-display-step.scss';

const TotalDisplayStep = ({vulnerabilityPerSeverity, cisDockerBenchmarkCountPerLevel, cisDockerBenchmarkScanEnabled}) => (
    <StepDisplay step="2"  title="Total vulnerabilities:" className="total-display-step">
        <VulnerabilitiesSummaryDisplay id="runtime-scan-vulnerabilities" vulnerabilities={vulnerabilityPerSeverity || []} />
        {!!cisDockerBenchmarkScanEnabled &&
            <div className="cis-benchmark-total-display">
                <StepDisplayTitle>CIS Docker Benchmark</StepDisplayTitle>
                <CisBenchmarkLevelsDisplay id="runtime-scan-cis" levels={cisDockerBenchmarkCountPerLevel || []} />
            </div>
        }
    </StepDisplay>
)

export default TotalDisplayStep;