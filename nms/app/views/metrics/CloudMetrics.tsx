/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type {MetricGraphConfig} from '../../components/insights/Metrics';
import type {TimeRange} from '../../components/insights/AsyncMetric';

import AppBar from '@mui/material/AppBar';
import AsyncMetric from '../../components/insights/AsyncMetric';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import ImageList from '@mui/material/ImageList';
import ImageListItem from '@mui/material/ImageListItem';
import React from 'react';
import Text from '../../theme/design-system/Text';
import TimeRangeSelector from '../../components/insights/TimeRangeSelector';

import {Theme} from '@mui/material/styles';
import {makeStyles} from '@mui/styles';
import {resolveQuery} from '../../components/insights/Metrics';
import {useState} from 'react';

const useStyles = makeStyles<Theme>(theme => ({
  formControl: {
    minWidth: '200px',
    padding: theme.spacing(),
  },
  appBar: {
    display: 'inline-block',
  },
}));

export default function CloudMetrics() {
  const classes = useStyles();
  const [timeRange, setTimeRange] = useState<TimeRange>('24_hours');

  const chartConfigs: Array<MetricGraphConfig> = [
    {
      label: 'REST API (2xx status code)',
      basicQueryConfigs: [],
      customQueryConfigs: [
        {
          resolveQuery: () => "sum(response_status{code=~'2..'})",
        },
      ],
      unit: '',
    },
    {
      label: 'REST API (3xx status code)',
      basicQueryConfigs: [],
      customQueryConfigs: [
        {
          resolveQuery: () => "sum(response_status{code=~'3..'})",
        },
      ],
      unit: '',
    },
    {
      label: 'REST API (4xx status code)',
      basicQueryConfigs: [],
      customQueryConfigs: [
        {
          resolveQuery: () => "sum(response_status{code=~'4..'})",
        },
      ],
      unit: '',
    },
    {
      label: 'REST API (5xx status code)',
      basicQueryConfigs: [],
      customQueryConfigs: [
        {
          resolveQuery: () => "sum(response_status{code=~'5..'})",
        },
      ],
      unit: '',
    },
  ];

  return (
    <>
      <AppBar className={classes.appBar} position="static" color="default">
        <TimeRangeSelector
          className={classes.formControl}
          value={timeRange}
          onChange={setTimeRange}
        />
      </AppBar>
      <ImageList cols={2} rowHeight={300}>
        {chartConfigs.map((config, i) => (
          <ImageListItem key={i} cols={1}>
            <Card>
              <CardContent>
                <Text variant="h6">{config.label}</Text>
                <div style={{height: 250}}>
                  <AsyncMetric
                    label={config.label}
                    unit={config.unit || ''}
                    queries={resolveQuery(config, 'gatewayID', '')}
                    timeRange={timeRange}
                    networkId="cloud"
                  />
                </div>
              </CardContent>
            </Card>
          </ImageListItem>
        ))}
      </ImageList>
    </>
  );
}
