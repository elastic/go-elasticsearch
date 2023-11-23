// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeAnalysisClassification type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L227-L236
type DataframeAnalysisClassification struct {
	// Alpha Advanced configuration option. Machine learning uses loss guided tree
	// growing, which means that the decision trees grow where the regularized loss
	// decreases most quickly. This parameter affects loss calculations by acting as
	// a multiplier of the tree depth. Higher alpha values result in shallower trees
	// and faster training times. By default, this value is calculated during
	// hyperparameter optimization. It must be greater than or equal to zero.
	Alpha                    *Float64 `json:"alpha,omitempty"`
	ClassAssignmentObjective *string  `json:"class_assignment_objective,omitempty"`
	// DependentVariable Defines which field of the document is to be predicted. It must match one of
	// the fields in the index being used to train. If this field is missing from a
	// document, then that document will not be used for training, but a prediction
	// with the trained model will be generated for it. It is also known as
	// continuous target variable.
	// For classification analysis, the data type of the field must be numeric
	// (`integer`, `short`, `long`, `byte`), categorical (`ip` or `keyword`), or
	// `boolean`. There must be no more than 30 different values in this field.
	// For regression analysis, the data type of the field must be numeric.
	DependentVariable string `json:"dependent_variable"`
	// DownsampleFactor Advanced configuration option. Controls the fraction of data that is used to
	// compute the derivatives of the loss function for tree training. A small value
	// results in the use of a small fraction of the data. If this value is set to
	// be less than 1, accuracy typically improves. However, too small a value may
	// result in poor convergence for the ensemble and so require more trees. By
	// default, this value is calculated during hyperparameter optimization. It must
	// be greater than zero and less than or equal to 1.
	DownsampleFactor *Float64 `json:"downsample_factor,omitempty"`
	// EarlyStoppingEnabled Advanced configuration option. Specifies whether the training process should
	// finish if it is not finding any better performing models. If disabled, the
	// training process can take significantly longer and the chance of finding a
	// better performing model is unremarkable.
	EarlyStoppingEnabled *bool `json:"early_stopping_enabled,omitempty"`
	// Eta Advanced configuration option. The shrinkage applied to the weights. Smaller
	// values result in larger forests which have a better generalization error.
	// However, larger forests cause slower training. By default, this value is
	// calculated during hyperparameter optimization. It must be a value between
	// 0.001 and 1.
	Eta *Float64 `json:"eta,omitempty"`
	// EtaGrowthRatePerTree Advanced configuration option. Specifies the rate at which `eta` increases
	// for each new tree that is added to the forest. For example, a rate of 1.05
	// increases `eta` by 5% for each extra tree. By default, this value is
	// calculated during hyperparameter optimization. It must be between 0.5 and 2.
	EtaGrowthRatePerTree *Float64 `json:"eta_growth_rate_per_tree,omitempty"`
	// FeatureBagFraction Advanced configuration option. Defines the fraction of features that will be
	// used when selecting a random bag for each candidate split. By default, this
	// value is calculated during hyperparameter optimization.
	FeatureBagFraction *Float64 `json:"feature_bag_fraction,omitempty"`
	// FeatureProcessors Advanced configuration option. A collection of feature preprocessors that
	// modify one or more included fields. The analysis uses the resulting one or
	// more features instead of the original document field. However, these features
	// are ephemeral; they are not stored in the destination index. Multiple
	// `feature_processors` entries can refer to the same document fields. Automatic
	// categorical feature encoding still occurs for the fields that are unprocessed
	// by a custom processor or that have categorical values. Use this property only
	// if you want to override the automatic feature encoding of the specified
	// fields.
	FeatureProcessors []DataframeAnalysisFeatureProcessor `json:"feature_processors,omitempty"`
	// Gamma Advanced configuration option. Regularization parameter to prevent
	// overfitting on the training data set. Multiplies a linear penalty associated
	// with the size of individual trees in the forest. A high gamma value causes
	// training to prefer small trees. A small gamma value results in larger
	// individual trees and slower training. By default, this value is calculated
	// during hyperparameter optimization. It must be a nonnegative value.
	Gamma *Float64 `json:"gamma,omitempty"`
	// Lambda Advanced configuration option. Regularization parameter to prevent
	// overfitting on the training data set. Multiplies an L2 regularization term
	// which applies to leaf weights of the individual trees in the forest. A high
	// lambda value causes training to favor small leaf weights. This behavior makes
	// the prediction function smoother at the expense of potentially not being able
	// to capture relevant relationships between the features and the dependent
	// variable. A small lambda value results in large individual trees and slower
	// training. By default, this value is calculated during hyperparameter
	// optimization. It must be a nonnegative value.
	Lambda *Float64 `json:"lambda,omitempty"`
	// MaxOptimizationRoundsPerHyperparameter Advanced configuration option. A multiplier responsible for determining the
	// maximum number of hyperparameter optimization steps in the Bayesian
	// optimization procedure. The maximum number of steps is determined based on
	// the number of undefined hyperparameters times the maximum optimization rounds
	// per hyperparameter. By default, this value is calculated during
	// hyperparameter optimization.
	MaxOptimizationRoundsPerHyperparameter *int `json:"max_optimization_rounds_per_hyperparameter,omitempty"`
	// MaxTrees Advanced configuration option. Defines the maximum number of decision trees
	// in the forest. The maximum value is 2000. By default, this value is
	// calculated during hyperparameter optimization.
	MaxTrees *int `json:"max_trees,omitempty"`
	// NumTopClasses Defines the number of categories for which the predicted probabilities are
	// reported. It must be non-negative or -1. If it is -1 or greater than the
	// total number of categories, probabilities are reported for all categories; if
	// you have a large number of categories, there could be a significant effect on
	// the size of your destination index. NOTE: To use the AUC ROC evaluation
	// method, `num_top_classes` must be set to -1 or a value greater than or equal
	// to the total number of categories.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// NumTopFeatureImportanceValues Advanced configuration option. Specifies the maximum number of feature
	// importance values per document to return. By default, no feature importance
	// calculation occurs.
	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`
	// PredictionFieldName Defines the name of the prediction field in the results. Defaults to
	// `<dependent_variable>_prediction`.
	PredictionFieldName *string `json:"prediction_field_name,omitempty"`
	// RandomizeSeed Defines the seed for the random generator that is used to pick training data.
	// By default, it is randomly generated. Set it to a specific value to use the
	// same training data each time you start a job (assuming other related
	// parameters such as `source` and `analyzed_fields` are the same).
	RandomizeSeed *Float64 `json:"randomize_seed,omitempty"`
	// SoftTreeDepthLimit Advanced configuration option. Machine learning uses loss guided tree
	// growing, which means that the decision trees grow where the regularized loss
	// decreases most quickly. This soft limit combines with the
	// `soft_tree_depth_tolerance` to penalize trees that exceed the specified
	// depth; the regularized loss increases quickly beyond this depth. By default,
	// this value is calculated during hyperparameter optimization. It must be
	// greater than or equal to 0.
	SoftTreeDepthLimit *int `json:"soft_tree_depth_limit,omitempty"`
	// SoftTreeDepthTolerance Advanced configuration option. This option controls how quickly the
	// regularized loss increases when the tree depth exceeds
	// `soft_tree_depth_limit`. By default, this value is calculated during
	// hyperparameter optimization. It must be greater than or equal to 0.01.
	SoftTreeDepthTolerance *Float64 `json:"soft_tree_depth_tolerance,omitempty"`
	// TrainingPercent Defines what percentage of the eligible documents that will be used for
	// training. Documents that are ignored by the analysis (for example those that
	// contain arrays with more than one value) won’t be included in the calculation
	// for used percentage.
	TrainingPercent Percentage `json:"training_percent,omitempty"`
}

func (s *DataframeAnalysisClassification) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "alpha":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Alpha = &f
			case float64:
				f := Float64(v)
				s.Alpha = &f
			}

		case "class_assignment_objective":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ClassAssignmentObjective = &o

		case "dependent_variable":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DependentVariable = o

		case "downsample_factor":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.DownsampleFactor = &f
			case float64:
				f := Float64(v)
				s.DownsampleFactor = &f
			}

		case "early_stopping_enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.EarlyStoppingEnabled = &value
			case bool:
				s.EarlyStoppingEnabled = &v
			}

		case "eta":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Eta = &f
			case float64:
				f := Float64(v)
				s.Eta = &f
			}

		case "eta_growth_rate_per_tree":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.EtaGrowthRatePerTree = &f
			case float64:
				f := Float64(v)
				s.EtaGrowthRatePerTree = &f
			}

		case "feature_bag_fraction":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.FeatureBagFraction = &f
			case float64:
				f := Float64(v)
				s.FeatureBagFraction = &f
			}

		case "feature_processors":
			if err := dec.Decode(&s.FeatureProcessors); err != nil {
				return err
			}

		case "gamma":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Gamma = &f
			case float64:
				f := Float64(v)
				s.Gamma = &f
			}

		case "lambda":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Lambda = &f
			case float64:
				f := Float64(v)
				s.Lambda = &f
			}

		case "max_optimization_rounds_per_hyperparameter":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOptimizationRoundsPerHyperparameter = &value
			case float64:
				f := int(v)
				s.MaxOptimizationRoundsPerHyperparameter = &f
			}

		case "max_trees", "maximum_number_trees":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxTrees = &value
			case float64:
				f := int(v)
				s.MaxTrees = &f
			}

		case "num_top_classes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopClasses = &value
			case float64:
				f := int(v)
				s.NumTopClasses = &f
			}

		case "num_top_feature_importance_values":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopFeatureImportanceValues = &value
			case float64:
				f := int(v)
				s.NumTopFeatureImportanceValues = &f
			}

		case "prediction_field_name":
			if err := dec.Decode(&s.PredictionFieldName); err != nil {
				return err
			}

		case "randomize_seed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.RandomizeSeed = &f
			case float64:
				f := Float64(v)
				s.RandomizeSeed = &f
			}

		case "soft_tree_depth_limit":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SoftTreeDepthLimit = &value
			case float64:
				f := int(v)
				s.SoftTreeDepthLimit = &f
			}

		case "soft_tree_depth_tolerance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.SoftTreeDepthTolerance = &f
			case float64:
				f := Float64(v)
				s.SoftTreeDepthTolerance = &f
			}

		case "training_percent":
			if err := dec.Decode(&s.TrainingPercent); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalysisClassification returns a DataframeAnalysisClassification.
func NewDataframeAnalysisClassification() *DataframeAnalysisClassification {
	r := &DataframeAnalysisClassification{}

	return r
}
