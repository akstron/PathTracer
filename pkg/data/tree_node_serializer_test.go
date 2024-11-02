package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataSerialization(t *testing.T) {
	root := &TreeNode{
		info: &DirNode{},
		children: []*TreeNode{
			{
				info: &FileNode{},
			},
			{
				info: &DirNode{},
				children: []*TreeNode{
					{
						info: &DirNode{},
						children: []*TreeNode{
							{
								info: &FileNode{},
							},
						},
					},
					{
						info: &DirNode{},
					},
					{
						info: &DirNode{},
					},
					{
						info: &FileNode{},
					},
				},
			},
			{
				info: &FileNode{},
			},
		},
	}

	serializedNode, err := json.Marshal(root)
	require.NoError(t, err)

	var extractedRoot TreeNode
	err = json.Unmarshal(serializedNode, &extractedRoot)
	require.NoError(t, err)

	trMg := NewTreeManager(&extractedRoot)
	it := NewTreeIterator(trMg)
	dirCnt := 0
	fileCnt := 0

	for it.HasNext() {
		got, err := it.Next()
		require.NoError(t, err)

		switch got.(type) {
		case *DirNode:
			dirCnt += 1
		case *FileNode:
			fileCnt += 1
		default:
			t.FailNow()
		}
	}

	require.Equal(t, 5, dirCnt)
	require.Equal(t, 4, fileCnt)
}
