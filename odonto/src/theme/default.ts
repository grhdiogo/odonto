import { colors } from './colors';
import { dimens } from './dimens';

const defaultTheme = {
  // colors
  fontColor: colors.textColor,
  fontNegative: colors.textNegative,
  fontDanger: colors.danger,
  fontSuccess: colors.textGreen,

  // background
  bgDefault: colors.bgColor,
  bgLight: colors.bgLight,
  bgGray: colors.bgGray,
  bgLightGray: colors.bgLightGray,
  bgBlue: colors.bgBlue,
  bgGreen: colors.bgGreen,

  // link
  linkMenu: colors.textNegative,
  linkMenuHover: colors.linkActive,
  linkActive: colors.linkActive,

  // buttons
  btnDanger: colors.danger,
  btnSuccess: colors.success,
  btnGosht: colors.btnGosht,
  btnInfo: colors.btnInfo,
  btnDefault: colors.btnDefault,

  // dimensions
  fontSize: dimens.textSizeNormal,
};

export default defaultTheme;
