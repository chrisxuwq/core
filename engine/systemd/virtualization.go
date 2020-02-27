package systemd

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/projecteru2/core/engine"
	enginetypes "github.com/projecteru2/core/engine/types"
	"github.com/projecteru2/core/utils"
)

const (
	cmdFileExist     = `/usr/bin/test -f '%s'`
	cmdCopyFromStdin = `/bin/cp -f /dev/stdin '%s'`
	cmdMkdir         = `/bin/mkdir -p %s`

	eruSystemdUnitPath = `/usr/local/lib/systemd/system/`
)

func (s *SystemdSSH) VirtualizationCreate(ctx context.Context, opts *enginetypes.VirtualizationCreateOptions) (created *enginetypes.VirtualizationCreated, err error) {
	CPUAmount, err := s.CPUInfo(ctx)
	if err != nil {
		return
	}

	buffer, err := s.newUnitBuilder(opts).buildUnit().buildResource(CPUAmount).buildExec().buffer()
	basename := fmt.Sprintf("%s.service", opts.Name)
	if err = s.VirtualizationCopyTo(ctx, "", filepath.Join(eruSystemdUnitPath, basename), buffer, true, true); err != nil {
		return
	}
	return &enginetypes.VirtualizationCreated{
		ID:   "SYSTEMD-" + utils.RandomString(46),
		Name: opts.Name,
	}, nil
}

func (s *SystemdSSH) VirtualizationCopyTo(ctx context.Context, ID, path string, content io.Reader, AllowOverwriteDirWithFile, _ bool) (err error) {
	// mkdir -p $(dirname $PATH)
	dirname, _ := filepath.Split(path)
	if _, stderr, err := s.runSingleCommand(ctx, fmt.Sprintf(cmdMkdir, dirname), nil); err != nil {
		return errors.Wrap(err, stderr.String())
	}

	// test -f $PATH && exit -1
	if !AllowOverwriteDirWithFile {
		if _, _, err = s.runSingleCommand(ctx, fmt.Sprintf(cmdFileExist, path), nil); err == nil {
			return fmt.Errorf("[VirtualizationCopyTo] file existed: %s", path)
		}
	}

	// cp /dev/stdin $PATH
	_, stderr, err := s.runSingleCommand(ctx, fmt.Sprintf(cmdCopyFromStdin, path), content)
	return errors.Wrap(err, stderr.String())
}

func (s *SystemdSSH) VirtualizationStart(ctx context.Context, ID string) (err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationStop(ctx context.Context, ID string) (err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationRemove(ctx context.Context, ID string, volumes, force bool) (err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationInspect(ctx context.Context, ID string) (info *enginetypes.VirtualizationInfo, err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationLogs(ctx context.Context, ID string, follow, stdout, stderr bool) (reader io.ReadCloser, err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationAttach(ctx context.Context, ID string, stream, stdin bool) (reader io.ReadCloser, writer io.WriteCloser, err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationResize(ctx context.Context, ID string, height, width uint) (err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationWait(ctx context.Context, ID, state string) (res *enginetypes.VirtualizationWaitResult, err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationUpdateResource(ctx context.Context, ID string, opts *enginetypes.VirtualizationResource) (err error) {
	err = engine.NotImplementedError
	return
}

func (s *SystemdSSH) VirtualizationCopyFrom(ctx context.Context, ID, path string) (reader io.ReadCloser, filename string, err error) {
	err = engine.NotImplementedError
	return
}
